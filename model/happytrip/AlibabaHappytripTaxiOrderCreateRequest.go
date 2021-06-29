package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户叫车 APIRequest
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

func NewAlibabaHappytripTaxiOrderCreateRequest() *AlibabaHappytripTaxiOrderCreateRequest{
    return &AlibabaHappytripTaxiOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.create"
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderCreateRequest) SetUid(uid string) error {
    r.uid = uid
    r.Set("uid", uid)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetUid() string {
    return r.uid
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetPassengerPhone(passengerPhone string) error {
    r.passengerPhone = passengerPhone
    r.Set("passenger_phone", passengerPhone)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetPassengerPhone() string {
    return r.passengerPhone
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCity(city string) error {
    r.city = city
    r.Set("city", city)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetCity() string {
    return r.city
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetFlat(flat string) error {
    r.flat = flat
    r.Set("flat", flat)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetFlat() string {
    return r.flat
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetFlng(flng string) error {
    r.flng = flng
    r.Set("flng", flng)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetFlng() string {
    return r.flng
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetStartName(startName string) error {
    r.startName = startName
    r.Set("start_name", startName)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetStartName() string {
    return r.startName
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetStartAddress(startAddress string) error {
    r.startAddress = startAddress
    r.Set("start_address", startAddress)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetStartAddress() string {
    return r.startAddress
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetTlat(tlat string) error {
    r.tlat = tlat
    r.Set("tlat", tlat)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetTlat() string {
    return r.tlat
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetTlng(tlng string) error {
    r.tlng = tlng
    r.Set("tlng", tlng)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetTlng() string {
    return r.tlng
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetEndName(endName string) error {
    r.endName = endName
    r.Set("end_name", endName)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetEndName() string {
    return r.endName
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetEndAddress(endAddress string) error {
    r.endAddress = endAddress
    r.Set("end_address", endAddress)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetEndAddress() string {
    return r.endAddress
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetClat(clat string) error {
    r.clat = clat
    r.Set("clat", clat)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetClat() string {
    return r.clat
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetClng(clng string) error {
    r.clng = clng
    r.Set("clng", clng)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetClng() string {
    return r.clng
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetDepartureTime(departureTime string) error {
    r.departureTime = departureTime
    r.Set("departure_time", departureTime)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetDepartureTime() string {
    return r.departureTime
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetRequireLevel(requireLevel string) error {
    r.requireLevel = requireLevel
    r.Set("require_level", requireLevel)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetRequireLevel() string {
    return r.requireLevel
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetAppTime(appTime string) error {
    r.appTime = appTime
    r.Set("app_time", appTime)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetAppTime() string {
    return r.appTime
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetMapType(mapType string) error {
    r.mapType = mapType
    r.Set("map_type", mapType)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetMapType() string {
    return r.mapType
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetSmsPolicy(smsPolicy int64) error {
    r.smsPolicy = smsPolicy
    r.Set("sms_policy", smsPolicy)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetSmsPolicy() int64 {
    return r.smsPolicy
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetExtraInfo(extraInfo string) error {
    r.extraInfo = extraInfo
    r.Set("extra_info", extraInfo)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetExtraInfo() string {
    return r.extraInfo
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetDynamicMd5(dynamicMd5 string) error {
    r.dynamicMd5 = dynamicMd5
    r.Set("dynamic_md5", dynamicMd5)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetDynamicMd5() string {
    return r.dynamicMd5
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCostCenter(costCenter string) error {
    r.costCenter = costCenter
    r.Set("cost_center", costCenter)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetCostCenter() string {
    return r.costCenter
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetLineType(lineType int64) error {
    r.lineType = lineType
    r.Set("line_type", lineType)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetLineType() int64 {
    return r.lineType
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCarpoolType(carpoolType int64) error {
    r.carpoolType = carpoolType
    r.Set("carpool_type", carpoolType)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetCarpoolType() int64 {
    return r.carpoolType
}

func (r *AlibabaHappytripTaxiOrderCreateRequest) SetPassengerNumber(passengerNumber int64) error {
    r.passengerNumber = passengerNumber
    r.Set("passenger_number", passengerNumber)
    return nil
}

func (r AlibabaHappytripTaxiOrderCreateRequest) GetPassengerNumber() int64 {
    return r.passengerNumber
}

