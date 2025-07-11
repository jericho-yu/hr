import { AnyArray } from "jericho-yu.aid"

export class EverydayModule {
  _data = AnyArray.make(0)

  constructor(data = []) { this._data = AnyArray.new(data) }

  static new = (data) => new EverydayModule(data)

  get data() { return this._data }

  find(date = '') {
    if (!date) return null

    this._data.copy().then(values => {
      if (values[6].value.match(/\b\d{4}-\d{2}-\d{2}\b/g)[0] || '' === date) {
        return values
      }
    })
  }
}
