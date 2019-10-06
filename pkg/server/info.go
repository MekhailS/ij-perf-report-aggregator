package server

import (
  "github.com/bvinc/go-sqlite-lite/sqlite3"
  "github.com/develar/errors"
  "net/http"
  "report-aggregator/pkg/util"
)

var essentialMetricNames = []string{"bootstrap", "appInitPreparation", "appInit", "pluginDescriptorLoading", "appComponentCreation", "projectComponentCreation"}
var instantMetricNames = []string{"splash", "startUpCompleted"}

func (t *StatsServer) handleInfoRequest(_ *http.Request) ([]byte, error) {
  productNames, err := t.getProductNames()
  if err != nil {
    return nil, err
  }

  statement, err := t.db.Prepare("select rowid as id, name from machine where rowid in (select distinct machine from report where product = ?) order by name")
  if err != nil {
    return nil, err
  }

  defer util.Close(statement, t.logger)

  var errRef error
  buffer := byteBufferPool.Get()
  defer byteBufferPool.Put(buffer)
  WriteInfo(buffer, productNames, essentialMetricNames, statement, &errRef)
  if errRef != nil {
    return nil, errRef
  }
  return CopyBuffer(buffer), nil
}

func (t *StatsServer) getProductNames() ([]string, error) {
  statement, err := t.db.Prepare("select distinct product from report order by product")
  if err != nil {
    return nil, errors.WithStack(err)
  }

  defer util.Close(statement, t.logger)

  return t.readStringList(statement)
}

func (t *StatsServer) readStringList(statement *sqlite3.Stmt) ([]string, error) {
  var result []string
  for {
    hasRow, err := statement.Step()
    if err != nil {
      return nil, errors.WithStack(err)
    }

    value, _, err := statement.ColumnText(0)
    if err != nil {
      return nil, errors.WithStack(err)
    }

    if !hasRow {
      break
    }

    result = append(result, value)
  }
  return result, nil
}