package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户叫车 API请求
alibaba.happytrip.taxi.order.create

用户根据需要发起叫车请求，在发起请求之前必须事先获得order id.
*/
type AlibabaHappytripTaxiOrderCreateRequest struct {
    model.Params
    // 用户唯一标识
    uid   string
    // 请求id 获取请参见
    orderId   string
    // 叫车车型，0(实时)；1(预约)
    type   int64
    // 乘客手机号
    passengerPhone   string
    // 出发地城市
    city   string
    // 出发地纬度
    flat   string
    // 出发地经度
    flng   string
    // 出发地名称(最多50个字)
    startName   string
    // 出发地详细地址(最多100个字)
    startAddress   string
    // 目的地纬度
    tlat   string
    // 目的地经度
    tlng   string
    // 目的地名称(最多50个字)
    endName   string
    // 目的地详细地址(最多100个字)
    endAddress   string
    // 当前位置纬度
    clat   string
    // 当前位置经度
    clng   string
    // 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
    departureTime   string
    // 车型代码
    requireLevel   string
    // 客户端时间（例如：2015-06-16 12:00:09）
    appTime   string
    // 地图类型:amap：高德，默认高德地图
    mapType   string
    // 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
    smsPolicy   int64
    // 备注
    extraInfo   string
    // 价格md5,通过 预估价接口获得
    dynamicMd5   string
    // 成本中心代码，用于区分不同的分账账号
    costCenter   string
    // 线路类型，0或空表示普通线路；1，表示一口价路线
    lineType   int64
    // 0：不拼车 1:允许拼车，默认不拼车
    carpoolType   int64
    // 乘车人数，允许拼车时该参数不为空
    passengerNumber   int64
}

// 初始化AlibabaHappytripTaxiOrderCreateRequest对象
func NewAlibabaHappytripTaxiOrderCreateRequest() *AlibabaHappytripTaxiOrderCreateRequest{
    return &AlibabaHappytripTaxiOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uid Setter
// 用户唯一标识
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetUid() string {
    return r.uid
}
// OrderId Setter
// 请求id 获取请参见
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetOrderId() string {
    return r.orderId
}
// Type Setter
// 叫车车型，0(实时)；1(预约)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetType() int64 {
    return r.type
}
// PassengerPhone Setter
// 乘客手机号
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetPassengerPhone(passengerPhone string) error {
    r.passengerPhone = passengerPhone
    r.Set("passenger_phone", passengerPhone)
    return nil
}

// PassengerPhone Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetPassengerPhone() string {
    return r.passengerPhone
}
// City Setter
// 出发地城市
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetCity() string {
    return r.city
}
// Flat Setter
// 出发地纬度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetFlat(flat string) error {
    r.flat = flat
    r.Set("flat", flat)
    return nil
}

// Flat Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetFlat() string {
    return r.flat
}
// Flng Setter
// 出发地经度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetFlng(flng string) error {
    r.flng = flng
    r.Set("flng", flng)
    return nil
}

// Flng Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetFlng() string {
    return r.flng
}
// StartName Setter
// 出发地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetStartName(startName string) error {
    r.startName = startName
    r.Set("start_name", startName)
    return nil
}

// StartName Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetStartName() string {
    return r.startName
}
// StartAddress Setter
// 出发地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetStartAddress(startAddress string) error {
    r.startAddress = startAddress
    r.Set("start_address", startAddress)
    return nil
}

// StartAddress Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetStartAddress() string {
    return r.startAddress
}
// Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetTlat(tlat string) error {
    r.tlat = tlat
    r.Set("tlat", tlat)
    return nil
}

// Tlat Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetTlat() string {
    return r.tlat
}
// Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetTlng(tlng string) error {
    r.tlng = tlng
    r.Set("tlng", tlng)
    return nil
}

// Tlng Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetTlng() string {
    return r.tlng
}
// EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetEndName(endName string) error {
    r.endName = endName
    r.Set("end_name", endName)
    return nil
}

// EndName Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetEndName() string {
    return r.endName
}
// EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetEndAddress(endAddress string) error {
    r.endAddress = endAddress
    r.Set("end_address", endAddress)
    return nil
}

// EndAddress Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetEndAddress() string {
    return r.endAddress
}
// Clat Setter
// 当前位置纬度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetClat(clat string) error {
    r.clat = clat
    r.Set("clat", clat)
    return nil
}

// Clat Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetClat() string {
    return r.clat
}
// Clng Setter
// 当前位置经度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetClng(clng string) error {
    r.clng = clng
    r.Set("clng", clng)
    return nil
}

// Clng Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetClng() string {
    return r.clng
}
// DepartureTime Setter
// 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetDepartureTime(departureTime string) error {
    r.departureTime = departureTime
    r.Set("departure_time", departureTime)
    return nil
}

// DepartureTime Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetDepartureTime() string {
    return r.departureTime
}
// RequireLevel Setter
// 车型代码
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetRequireLevel(requireLevel string) error {
    r.requireLevel = requireLevel
    r.Set("require_level", requireLevel)
    return nil
}

// RequireLevel Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetRequireLevel() string {
    return r.requireLevel
}
// AppTime Setter
// 客户端时间（例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetAppTime(appTime string) error {
    r.appTime = appTime
    r.Set("app_time", appTime)
    return nil
}

// AppTime Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetAppTime() string {
    return r.appTime
}
// MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetMapType(mapType string) error {
    r.mapType = mapType
    r.Set("map_type", mapType)
    return nil
}

// MapType Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetMapType() string {
    return r.mapType
}
// SmsPolicy Setter
// 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetSmsPolicy(smsPolicy int64) error {
    r.smsPolicy = smsPolicy
    r.Set("sms_policy", smsPolicy)
    return nil
}

// SmsPolicy Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetSmsPolicy() int64 {
    return r.smsPolicy
}
// ExtraInfo Setter
// 备注
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetExtraInfo(extraInfo string) error {
    r.extraInfo = extraInfo
    r.Set("extra_info", extraInfo)
    return nil
}

// ExtraInfo Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetExtraInfo() string {
    return r.extraInfo
}
// DynamicMd5 Setter
// 价格md5,通过 预估价接口获得
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetDynamicMd5(dynamicMd5 string) error {
    r.dynamicMd5 = dynamicMd5
    r.Set("dynamic_md5", dynamicMd5)
    return nil
}

// DynamicMd5 Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetDynamicMd5() string {
    return r.dynamicMd5
}
// CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCostCenter(costCenter string) error {
    r.costCenter = costCenter
    r.Set("cost_center", costCenter)
    return nil
}

// CostCenter Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetCostCenter() string {
    return r.costCenter
}
// LineType Setter
// 线路类型，0或空表示普通线路；1，表示一口价路线
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetLineType(lineType int64) error {
    r.lineType = lineType
    r.Set("line_type", lineType)
    return nil
}

// LineType Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetLineType() int64 {
    return r.lineType
}
// CarpoolType Setter
// 0：不拼车 1:允许拼车，默认不拼车
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCarpoolType(carpoolType int64) error {
    r.carpoolType = carpoolType
    r.Set("carpool_type", carpoolType)
    return nil
}

// CarpoolType Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetCarpoolType() int64 {
    return r.carpoolType
}
// PassengerNumber Setter
// 乘车人数，允许拼车时该参数不为空
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetPassengerNumber(passengerNumber int64) error {
    r.passengerNumber = passengerNumber
    r.Set("passenger_number", passengerNumber)
    return nil
}

// PassengerNumber Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetPassengerNumber() int64 {
    return r.passengerNumber
}
