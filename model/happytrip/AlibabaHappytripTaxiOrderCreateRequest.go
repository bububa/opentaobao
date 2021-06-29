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
    _uid   string
    // 请求id 获取请参见
    _orderId   string
    // 叫车车型，0(实时)；1(预约)
    _type   int64
    // 乘客手机号
    _passengerPhone   string
    // 出发地城市
    _city   string
    // 出发地纬度
    _flat   string
    // 出发地经度
    _flng   string
    // 出发地名称(最多50个字)
    _startName   string
    // 出发地详细地址(最多100个字)
    _startAddress   string
    // 目的地纬度
    _tlat   string
    // 目的地经度
    _tlng   string
    // 目的地名称(最多50个字)
    _endName   string
    // 目的地详细地址(最多100个字)
    _endAddress   string
    // 当前位置纬度
    _clat   string
    // 当前位置经度
    _clng   string
    // 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
    _departureTime   string
    // 车型代码
    _requireLevel   string
    // 客户端时间（例如：2015-06-16 12:00:09）
    _appTime   string
    // 地图类型:amap：高德，默认高德地图
    _mapType   string
    // 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
    _smsPolicy   int64
    // 备注
    _extraInfo   string
    // 价格md5,通过 预估价接口获得
    _dynamicMd5   string
    // 成本中心代码，用于区分不同的分账账号
    _costCenter   string
    // 线路类型，0或空表示普通线路；1，表示一口价路线
    _lineType   int64
    // 0：不拼车 1:允许拼车，默认不拼车
    _carpoolType   int64
    // 乘车人数，允许拼车时该参数不为空
    _passengerNumber   int64
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
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetUid(_uid string) error {
    r._uid = _uid
    r.Set("uid", _uid)
    return nil
}

// Uid Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetUid() string {
    return r._uid
}
// OrderId Setter
// 请求id 获取请参见
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetOrderId() string {
    return r._orderId
}
// Type Setter
// 叫车车型，0(实时)；1(预约)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetType() int64 {
    return r._type
}
// PassengerPhone Setter
// 乘客手机号
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetPassengerPhone(_passengerPhone string) error {
    r._passengerPhone = _passengerPhone
    r.Set("passenger_phone", _passengerPhone)
    return nil
}

// PassengerPhone Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetPassengerPhone() string {
    return r._passengerPhone
}
// City Setter
// 出发地城市
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCity(_city string) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetCity() string {
    return r._city
}
// Flat Setter
// 出发地纬度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetFlat(_flat string) error {
    r._flat = _flat
    r.Set("flat", _flat)
    return nil
}

// Flat Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetFlat() string {
    return r._flat
}
// Flng Setter
// 出发地经度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetFlng(_flng string) error {
    r._flng = _flng
    r.Set("flng", _flng)
    return nil
}

// Flng Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetFlng() string {
    return r._flng
}
// StartName Setter
// 出发地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetStartName(_startName string) error {
    r._startName = _startName
    r.Set("start_name", _startName)
    return nil
}

// StartName Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetStartName() string {
    return r._startName
}
// StartAddress Setter
// 出发地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetStartAddress(_startAddress string) error {
    r._startAddress = _startAddress
    r.Set("start_address", _startAddress)
    return nil
}

// StartAddress Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetStartAddress() string {
    return r._startAddress
}
// Tlat Setter
// 目的地纬度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetTlat(_tlat string) error {
    r._tlat = _tlat
    r.Set("tlat", _tlat)
    return nil
}

// Tlat Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetTlat() string {
    return r._tlat
}
// Tlng Setter
// 目的地经度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetTlng(_tlng string) error {
    r._tlng = _tlng
    r.Set("tlng", _tlng)
    return nil
}

// Tlng Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetTlng() string {
    return r._tlng
}
// EndName Setter
// 目的地名称(最多50个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetEndName(_endName string) error {
    r._endName = _endName
    r.Set("end_name", _endName)
    return nil
}

// EndName Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetEndName() string {
    return r._endName
}
// EndAddress Setter
// 目的地详细地址(最多100个字)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetEndAddress(_endAddress string) error {
    r._endAddress = _endAddress
    r.Set("end_address", _endAddress)
    return nil
}

// EndAddress Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetEndAddress() string {
    return r._endAddress
}
// Clat Setter
// 当前位置纬度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetClat(_clat string) error {
    r._clat = _clat
    r.Set("clat", _clat)
    return nil
}

// Clat Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetClat() string {
    return r._clat
}
// Clng Setter
// 当前位置经度
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetClng(_clng string) error {
    r._clng = _clng
    r.Set("clng", _clng)
    return nil
}

// Clng Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetClng() string {
    return r._clng
}
// DepartureTime Setter
// 出发时间，不传表示现在用车（例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetDepartureTime(_departureTime string) error {
    r._departureTime = _departureTime
    r.Set("departure_time", _departureTime)
    return nil
}

// DepartureTime Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetDepartureTime() string {
    return r._departureTime
}
// RequireLevel Setter
// 车型代码
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetRequireLevel(_requireLevel string) error {
    r._requireLevel = _requireLevel
    r.Set("require_level", _requireLevel)
    return nil
}

// RequireLevel Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetRequireLevel() string {
    return r._requireLevel
}
// AppTime Setter
// 客户端时间（例如：2015-06-16 12:00:09）
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetAppTime(_appTime string) error {
    r._appTime = _appTime
    r.Set("app_time", _appTime)
    return nil
}

// AppTime Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetAppTime() string {
    return r._appTime
}
// MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetMapType(_mapType string) error {
    r._mapType = _mapType
    r.Set("map_type", _mapType)
    return nil
}

// MapType Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetMapType() string {
    return r._mapType
}
// SmsPolicy Setter
// 发送短信策略(0,为叫车人和乘车人都发送，1，乘车人发送叫车人不发，2乘车人不发叫车人发，3乘车人和叫车人都不发)
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetSmsPolicy(_smsPolicy int64) error {
    r._smsPolicy = _smsPolicy
    r.Set("sms_policy", _smsPolicy)
    return nil
}

// SmsPolicy Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetSmsPolicy() int64 {
    return r._smsPolicy
}
// ExtraInfo Setter
// 备注
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetExtraInfo(_extraInfo string) error {
    r._extraInfo = _extraInfo
    r.Set("extra_info", _extraInfo)
    return nil
}

// ExtraInfo Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetExtraInfo() string {
    return r._extraInfo
}
// DynamicMd5 Setter
// 价格md5,通过 预估价接口获得
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetDynamicMd5(_dynamicMd5 string) error {
    r._dynamicMd5 = _dynamicMd5
    r.Set("dynamic_md5", _dynamicMd5)
    return nil
}

// DynamicMd5 Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetDynamicMd5() string {
    return r._dynamicMd5
}
// CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCostCenter(_costCenter string) error {
    r._costCenter = _costCenter
    r.Set("cost_center", _costCenter)
    return nil
}

// CostCenter Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetCostCenter() string {
    return r._costCenter
}
// LineType Setter
// 线路类型，0或空表示普通线路；1，表示一口价路线
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetLineType(_lineType int64) error {
    r._lineType = _lineType
    r.Set("line_type", _lineType)
    return nil
}

// LineType Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetLineType() int64 {
    return r._lineType
}
// CarpoolType Setter
// 0：不拼车 1:允许拼车，默认不拼车
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetCarpoolType(_carpoolType int64) error {
    r._carpoolType = _carpoolType
    r.Set("carpool_type", _carpoolType)
    return nil
}

// CarpoolType Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetCarpoolType() int64 {
    return r._carpoolType
}
// PassengerNumber Setter
// 乘车人数，允许拼车时该参数不为空
func (r *AlibabaHappytripTaxiOrderCreateRequest) SetPassengerNumber(_passengerNumber int64) error {
    r._passengerNumber = _passengerNumber
    r.Set("passenger_number", _passengerNumber)
    return nil
}

// PassengerNumber Getter
func (r AlibabaHappytripTaxiOrderCreateRequest) GetPassengerNumber() int64 {
    return r._passengerNumber
}
