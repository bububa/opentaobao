package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取价格预估信息 API请求
alibaba.happytrip.taxi.price.get

打车价格预估
*/
type AlibabaHappytripTaxiPriceGetAPIRequest struct {
    model.Params
    // 出发地纬度
    _flat   string
    // 出发地经度
    _flng   string
    // 目的地纬度
    _tlat   string
    // 目的地经度
    _tlng   string
    // 地图类型:amap：高德，默认高德地图
    _mapType   string
    // 出发城市id
    _city   string
    // 0:实时单 1:预约单
    _type   int64
    // 预约单必须传（格式例如：2015-06-16 12:00:09）
    _departureTime   string
    // 成本中心代码，用于区分不同的分账账号
    _costCenter   string
    // 供应商车型代码
    _requireLevel   string
    // 0：不拼车 1:允许拼车，默认不拼车
    _carpoolType   int64
    // 乘车人数
    _passengerNumber   int64
    // 用户唯一标识
    _uid   string
    // 乘客手机号
    _passengerPhone   string
}

// 初始化AlibabaHappytripTaxiPriceGetAPIRequest对象
func NewAlibabaHappytripTaxiPriceGetRequest() *AlibabaHappytripTaxiPriceGetAPIRequest{
    return &AlibabaHappytripTaxiPriceGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.price.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Flat Setter
// 出发地纬度
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetFlat(_flat string) error {
    r._flat = _flat
    r.Set("flat", _flat)
    return nil
}

// Flat Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetFlat() string {
    return r._flat
}
// Flng Setter
// 出发地经度
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetFlng(_flng string) error {
    r._flng = _flng
    r.Set("flng", _flng)
    return nil
}

// Flng Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetFlng() string {
    return r._flng
}
// Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetTlat(_tlat string) error {
    r._tlat = _tlat
    r.Set("tlat", _tlat)
    return nil
}

// Tlat Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetTlat() string {
    return r._tlat
}
// Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetTlng(_tlng string) error {
    r._tlng = _tlng
    r.Set("tlng", _tlng)
    return nil
}

// Tlng Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetTlng() string {
    return r._tlng
}
// MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetMapType(_mapType string) error {
    r._mapType = _mapType
    r.Set("map_type", _mapType)
    return nil
}

// MapType Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetMapType() string {
    return r._mapType
}
// City Setter
// 出发城市id
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetCity() string {
    return r._city
}
// Type Setter
// 0:实时单 1:预约单
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetType() int64 {
    return r._type
}
// DepartureTime Setter
// 预约单必须传（格式例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetDepartureTime(_departureTime string) error {
    r._departureTime = _departureTime
    r.Set("departure_time", _departureTime)
    return nil
}

// DepartureTime Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetDepartureTime() string {
    return r._departureTime
}
// CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetCostCenter(_costCenter string) error {
    r._costCenter = _costCenter
    r.Set("cost_center", _costCenter)
    return nil
}

// CostCenter Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetCostCenter() string {
    return r._costCenter
}
// RequireLevel Setter
// 供应商车型代码
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetRequireLevel(_requireLevel string) error {
    r._requireLevel = _requireLevel
    r.Set("require_level", _requireLevel)
    return nil
}

// RequireLevel Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetRequireLevel() string {
    return r._requireLevel
}
// CarpoolType Setter
// 0：不拼车 1:允许拼车，默认不拼车
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetCarpoolType(_carpoolType int64) error {
    r._carpoolType = _carpoolType
    r.Set("carpool_type", _carpoolType)
    return nil
}

// CarpoolType Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetCarpoolType() int64 {
    return r._carpoolType
}
// PassengerNumber Setter
// 乘车人数
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetPassengerNumber(_passengerNumber int64) error {
    r._passengerNumber = _passengerNumber
    r.Set("passenger_number", _passengerNumber)
    return nil
}

// PassengerNumber Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetPassengerNumber() int64 {
    return r._passengerNumber
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetUid(_uid string) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetUid() string {
    return r._uid
}
// PassengerPhone Setter
// 乘客手机号
func (r *AlibabaHappytripTaxiPriceGetAPIRequest) SetPassengerPhone(_passengerPhone string) error {
    r._passengerPhone = _passengerPhone
    r.Set("passenger_phone", _passengerPhone)
    return nil
}

// PassengerPhone Getter
func (r AlibabaHappytripTaxiPriceGetAPIRequest) GetPassengerPhone() string {
    return r._passengerPhone
}
