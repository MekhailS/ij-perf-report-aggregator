import { Ref } from "vue"
import { ServerConfigurator } from "./configurators/ServerConfigurator"
import { TimeRange } from "./configurators/TimeRangeConfigurator"
import { encodeRison } from "./rison"

export class Accident {
  constructor(readonly affectedTest: string, readonly date: string, readonly reason: string, readonly id: number, readonly buildNumber: string) {}
}

export function isDateInsideRange(dateOfAccident: Date, interval: TimeRange): boolean {
  const currentDate = new Date()
  const day = 24 * 60 * 60 * 1000
  const intervalMapping = {
    "1w": day * 7,
    "1M": day * 30,
    "3M": day * 30 * 3,
    "1y": day * 365,
    "all": day * 365,
  }
  const selectedDate = new Date()
  selectedDate.setTime(Date.now() - intervalMapping[interval])
  return dateOfAccident >= selectedDate && dateOfAccident <= currentDate
}

export function getWarningFromMetaDb(warnings: Ref<Array<Accident> | undefined>,
                                     branches: Array<string> | string | null,
                                     tests: Array<string> | string | null,
                                     table: string,
                                     timeRange: TimeRange) {
  if (branches == null) {
    return
  }
  if (!Array.isArray(branches)) {
    branches = [branches]
  }
  if (tests != null && !Array.isArray(tests)) {
    tests = [tests]
  }
  const url = ServerConfigurator.DEFAULT_SERVER_URL + "/api/meta/"
  warnings.value = []
  const data = tests == null ? {table, branches} : {table, branches, tests}
  fetch(url + encodeRison(data))
    .then(response => response.json())
    .then((data: Array<Accident>) => {
      for (const datum of data) {
        if (isDateInsideRange(new Date(datum.date), timeRange)) {
          warnings.value?.push(datum)
        }
      }
    })
    .catch(error => console.error(error))
}

export function isValueShouldBeMarked(accidents: Array<Accident> | null, value: Array<string>): boolean {
  return getAccident(accidents, value) != null
}

export function getAccident(accidents: Array<Accident> | null, value: Array<string>): Accident | null {
  if (accidents != null) {
    //perf db
    if (value.length == 10) {
      return accidents.find(accident => accident.affectedTest == value[5] && accident.buildNumber == value[7] + "." + value[8]) ?? null
    }
    //perf dev db
    if (value.length == 6) {
      return accidents.find(accident => accident.affectedTest == value[5] && accident.buildNumber == value[4]) ?? null
    }
  }
  return null
}