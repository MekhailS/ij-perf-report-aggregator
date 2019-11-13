package main

import (
  "flag"
  "fmt"
  "github.com/AlexAkulov/clickhouse-backup/pkg/chbackup"
  "github.com/JetBrains/ij-perf-report-aggregator/pkg/util"
  "github.com/develar/errors"
  "github.com/kelseyhightower/envconfig"
  "github.com/robfig/cron/v3"
  "go.uber.org/zap"
  "strings"
  "time"
)

var lastLocalBackups []string
var lastRemoteBackups []string
const backupSuffix = ".tar"

var incrementalBackupCounter = 0

func main() {
  logger := util.CreateLogger()
  defer func() {
    _ = logger.Sync()
  }()

  spec := flag.String("spec", "", "")
  flag.Parse()

  err := schedule(*spec, logger)
  if err != nil {
    logger.Fatal(fmt.Sprintf("%+v", err))
  }
}

func schedule(spec string, logger *zap.Logger) error {
  scheduler := cron.New(cron.WithLogger(&ZapLoggerAdapter{logger: logger}))

  config := chbackup.DefaultConfig()
  config.General.DisableProgressBar = true
  err := envconfig.Process("", config)
  if err != nil {
    return errors.WithStack(err)
  }

  _, err = scheduler.AddFunc(spec, func() {
    err := backup(config, logger)
    if err != nil {
      logger.Error("cannot backup", zap.Error(err))
    }
  })
  if err != nil {
    return err
  }

  scheduler.Run()
  return nil
}

func backup(config *chbackup.Config, logger *zap.Logger) error {
  backupName := time.Now().Format(chbackup.BackupTimeFormat)

  err := chbackup.CreateBackup(*config, backupName, "")
  if err != nil {
    return errors.WithStack(err)
  }

  var diffFrom string
  // first or each 4 backup is full
  if len(lastLocalBackups) < 1 || incrementalBackupCounter > 4 {
    incrementalBackupCounter = 0
    diffFrom = ""
  } else {
    diffFrom = strings.TrimSuffix(lastLocalBackups[len(lastLocalBackups) - 1], backupSuffix)
  }

  err = chbackup.Upload(*config, backupName, diffFrom)
  if err != nil {
    return errors.WithStack(err)
  }

  incrementalBackupCounter++

  if len(lastLocalBackups) > 2 {
    outdatedBackupName := lastLocalBackups[0]
    logger.Info("remove local backup", zap.String("name", outdatedBackupName))
    err := chbackup.RemoveBackupLocal(*config, outdatedBackupName)
    if err != nil {
      return errors.WithStack(err)
    }

    lastLocalBackups = append(lastLocalBackups[1:], backupName+backupSuffix)
  } else {
    lastLocalBackups = append(lastLocalBackups, backupName+backupSuffix)
  }

  if len(lastRemoteBackups) > 3 {
    outdatedBackupName := lastRemoteBackups[0]
    logger.Info("remove remote backup", zap.String("name", outdatedBackupName))
    err := chbackup.RemoveBackupRemote(*config, outdatedBackupName)
    if err != nil {
      return errors.WithStack(err)
    }

    lastRemoteBackups = append(lastRemoteBackups[1:], backupName + backupSuffix)
  } else {
    lastRemoteBackups = append(lastRemoteBackups, backupName + backupSuffix)
  }

  return nil
}

type ZapLoggerAdapter struct {
  logger *zap.Logger
}

func toZapFields(keysAndValues []interface{}) []zap.Field {
  var fields []zap.Field
  for i := 0; i < len(keysAndValues); i += 2 {
    key := keysAndValues[i].(string)
    value := keysAndValues[i+1]
    switch v := (value.(interface{})).(type) {
    case time.Time:
      fields = append(fields, zap.Time(key, v))
    case string:
      fields = append(fields, zap.String(key, v))
    case int:
      fields = append(fields, zap.Int(key, v))
    default:
      fields = append(fields, zap.String(key, fmt.Sprintf("%v", value)))
    }
  }
  return fields
}

func (t *ZapLoggerAdapter) Info(msg string, keysAndValues ...interface{}) {
  t.logger.Info(msg, toZapFields(keysAndValues)...)
}

func (t *ZapLoggerAdapter) Error(err error, msg string, keysAndValues ...interface{}) {
  var fields []zap.Field
  fields = append(fields, zap.Error(err))
  fields = append(fields, toZapFields(keysAndValues)...)
  t.logger.Error(msg, fields...)
}
