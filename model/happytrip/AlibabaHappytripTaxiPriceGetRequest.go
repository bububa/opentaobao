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
type AlibabaHappytripTaxiPriceGetRequest struct {
    model.Params
    // 出发地纬度
    flat   string
    // 出发地经度
    flng   string
    // 目的地纬度
    tlat   string
    // 目的地经度
    tlng   string
    // 地图类型:amap：高德，默认高德地图
    mapType   string
    // 出发城市id
    city   string
    // 0:实时单 1:预约单
    type   int64
    // 预约单必须传（格式例如：2015-06-16 12:00:09）
    departureTime   string
    // 成本中心代码，用于区分不同的分账账号
    costCenter   string
    // 供应商车型代码
    requireLevel   string
    // 0：不拼车 1:允许拼车，默认不拼车
    carpoolType   int64
    // 乘车人数
    passengerNumber   int64
    // 用户唯一标识
    uid   string
    // 乘客手机号
    passengerPhone   string
}

// 初始化AlibabaHappytripTaxiPriceGetRequest对象
func NewAlibabaHappytripTaxiPriceGetRequest() *AlibabaHappytripTaxiPriceGetRequest{
    return &AlibabaHappytripTaxiPriceGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiPriceGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.price.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiPriceGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Flat Setter
// 出发地纬度
func (r *AlibabaHappytripTaxiPriceGetRequest) SetFlat(flat string) error {
    r.flat = flat
    r.Set("flat", flat)
    return nil
}

// Flat Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetFlat() string {
    return r.flat
}
// Flng Setter
// 出发地经度
func (r *AlibabaHappytripTaxiPriceGetRequest) SetFlng(flng string) error {
    r.flng = flng
    r.Set("flng", flng)
    return nil
}

// Flng Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetFlng() string {
    return r.flng
}
// Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiPriceGetRequest) SetTlat(tlat string) error {
    r.tlat = tlat
    r.Set("tlat", tlat)
    return nil
}

// Tlat Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetTlat() string {
    return r.tlat
}
// Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiPriceGetRequest) SetTlng(tlng string) error {
    r.tlng = tlng
    r.Set("tlng", tlng)
    return nil
}

// Tlng Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetTlng() string {
    return r.tlng
}
// MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiPriceGetRequest) SetMapType(mapType string) error {
    r.mapType = mapType
    r.Set("map_type", mapType)
    return nil
}

// MapType Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetMapType() string {
    return r.mapType
}
// City Setter
// 出发城市id
func (r *AlibabaHappytripTaxiPriceGetRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetCity() string {
    return r.city
}
// Type Setter
// 0:实时单 1:预约单
func (r *AlibabaHappytripTaxiPriceGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetType() int64 {
    return r.type
}
// DepartureTime Setter
// 预约单必须传（格式例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiPriceGetRequest) SetDepartureTime(departureTime string) error {
    r.departureTime = departureTime
    r.Set("departure_time", departureTime)
    return nil
}

// DepartureTime Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetDepartureTime() string {
    return r.departureTime
}
// CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiPriceGetRequest) SetCostCenter(costCenter string) error {
    r.costCenter = costCenter
    r.Set("cost_center", costCenter)
    return nil
}

// CostCenter Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetCostCenter() string {
    return r.costCenter
}
// RequireLevel Setter
// 供应商车型代码
func (r *AlibabaHappytripTaxiPriceGetRequest) SetRequireLevel(requireLevel string) error {
    r.requireLevel = requireLevel
    r.Set("require_level", requireLevel)
    return nil
}

// RequireLevel Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetRequireLevel() string {
    return r.requireLevel
}
// CarpoolType Setter
// 0：不拼车 1:允许拼车，默认不拼车
func (r *AlibabaHappytripTaxiPriceGetRequest) SetCarpoolType(carpoolType int64) error {
    r.carpoolType = carpoolType
    r.Set("carpool_type", carpoolType)
    return nil
}

// CarpoolType Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetCarpoolType() int64 {
    return r.carpoolType
}
// PassengerNumber Setter
// 乘车人数
func (r *AlibabaHappytripTaxiPriceGetRequest) SetPassengerNumber(passengerNumber int64) error {
    r.passengerNumber = passengerNumber
    r.Set("passenger_number", passengerNumber)
    return nil
}

// PassengerNumber Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetPassengerNumber() int64 {
    return r.passengerNumber
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiPriceGetRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetUid() string {
    return r.uid
}
// PassengerPhone Setter
// 乘客手机号
func (r *AlibabaHappytripTaxiPriceGetRequest) SetPassengerPhone(passengerPhone string) error {
    r.passengerPhone = passengerPhone
    r.Set("passenger_phone", passengerPhone)
    return nil
}

// PassengerPhone Getter
func (r AlibabaHappytripTaxiPriceGetRequest) GetPassengerPhone() string {
    return r.passengerPhone
}
