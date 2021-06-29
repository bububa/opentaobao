package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专享房价 API请求
taobao.xhotel.rate.add

酒店产品库rate添加
*/
type TaobaoXhotelRateAddRequest struct {
    model.Params
    // gid酒店商品id
    _gid   int64
    // 酒店RPID
    _rpid   int64
    // 名称
    _name   string
    // 额外服务-是否可以加床，1：不可以，2：可以
    _addBed   int64
    // 额外服务-加床价格
    _addBedPrice   int64
    // 币种（仅支持CNY）
    _currencyCode   int64
    // 实价有房标签（RP支付类型为全额支付）
    _shijiaTag   int64
    // 价格和库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
    _inventoryPrice   string
    // “即时确认”标识，此类商品预订后直接发货。
    _jishiquerenTag   int64
    // 卖家自己系统的Code，简称RateCode
    _rateplanCode   string
    // 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
    _vendor   string
    // 卖家房型ID, 这是卖家自己系统中的房型ID，注意：需按照规则组合
    _outRid   string
    // 在添加新rate时，同时添加rate开关日历。可以只设定想设定的某些天，可以不连续。date：开关状态控制的是那一天rate_status：开关状态。0，关闭；1，打开
    _rateSwitchCal   string
    // 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    _lockEndTime   string
    // 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    _lockStartTime   string
    // 币种信息,默认是CNY,  @see com.taobao.trip.hotel.model.enums.CurrencyEnum
    _currencyCodeName   string
    // 操作人信息
    _operator   string
    // 默认是2 ,
    _source   int64
    // 1是开,0是关, 不填默认是开, rate状态
    _status   int64
    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    _onlineBookingBindingInfo   string
}

// 初始化TaobaoXhotelRateAddRequest对象
func NewTaobaoXhotelRateAddRequest() *TaobaoXhotelRateAddRequest{
    return &TaobaoXhotelRateAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Gid Setter
// gid酒店商品id
func (r *TaobaoXhotelRateAddRequest) SetGid(_gid int64) error {
    r._gid = _gid
    r.Set("gid", _gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelRateAddRequest) GetGid() int64 {
    return r._gid
}
// Rpid Setter
// 酒店RPID
func (r *TaobaoXhotelRateAddRequest) SetRpid(_rpid int64) error {
    r._rpid = _rpid
    r.Set("rpid", _rpid)
    return nil
}

// Rpid Getter
func (r TaobaoXhotelRateAddRequest) GetRpid() int64 {
    return r._rpid
}
// Name Setter
// 名称
func (r *TaobaoXhotelRateAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoXhotelRateAddRequest) GetName() string {
    return r._name
}
// AddBed Setter
// 额外服务-是否可以加床，1：不可以，2：可以
func (r *TaobaoXhotelRateAddRequest) SetAddBed(_addBed int64) error {
    r._addBed = _addBed
    r.Set("add_bed", _addBed)
    return nil
}

// AddBed Getter
func (r TaobaoXhotelRateAddRequest) GetAddBed() int64 {
    return r._addBed
}
// AddBedPrice Setter
// 额外服务-加床价格
func (r *TaobaoXhotelRateAddRequest) SetAddBedPrice(_addBedPrice int64) error {
    r._addBedPrice = _addBedPrice
    r.Set("add_bed_price", _addBedPrice)
    return nil
}

// AddBedPrice Getter
func (r TaobaoXhotelRateAddRequest) GetAddBedPrice() int64 {
    return r._addBedPrice
}
// CurrencyCode Setter
// 币种（仅支持CNY）
func (r *TaobaoXhotelRateAddRequest) SetCurrencyCode(_currencyCode int64) error {
    r._currencyCode = _currencyCode
    r.Set("currency_code", _currencyCode)
    return nil
}

// CurrencyCode Getter
func (r TaobaoXhotelRateAddRequest) GetCurrencyCode() int64 {
    return r._currencyCode
}
// ShijiaTag Setter
// 实价有房标签（RP支付类型为全额支付）
func (r *TaobaoXhotelRateAddRequest) SetShijiaTag(_shijiaTag int64) error {
    r._shijiaTag = _shijiaTag
    r.Set("shijia_tag", _shijiaTag)
    return nil
}

// ShijiaTag Getter
func (r TaobaoXhotelRateAddRequest) GetShijiaTag() int64 {
    return r._shijiaTag
}
// InventoryPrice Setter
// 价格和库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
func (r *TaobaoXhotelRateAddRequest) SetInventoryPrice(_inventoryPrice string) error {
    r._inventoryPrice = _inventoryPrice
    r.Set("inventory_price", _inventoryPrice)
    return nil
}

// InventoryPrice Getter
func (r TaobaoXhotelRateAddRequest) GetInventoryPrice() string {
    return r._inventoryPrice
}
// JishiquerenTag Setter
// “即时确认”标识，此类商品预订后直接发货。
func (r *TaobaoXhotelRateAddRequest) SetJishiquerenTag(_jishiquerenTag int64) error {
    r._jishiquerenTag = _jishiquerenTag
    r.Set("jishiqueren_tag", _jishiquerenTag)
    return nil
}

// JishiquerenTag Getter
func (r TaobaoXhotelRateAddRequest) GetJishiquerenTag() int64 {
    return r._jishiquerenTag
}
// RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoXhotelRateAddRequest) SetRateplanCode(_rateplanCode string) error {
    r._rateplanCode = _rateplanCode
    r.Set("rateplan_code", _rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateAddRequest) GetRateplanCode() string {
    return r._rateplanCode
}
// Vendor Setter
// 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
func (r *TaobaoXhotelRateAddRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateAddRequest) GetVendor() string {
    return r._vendor
}
// OutRid Setter
// 卖家房型ID, 这是卖家自己系统中的房型ID，注意：需按照规则组合
func (r *TaobaoXhotelRateAddRequest) SetOutRid(_outRid string) error {
    r._outRid = _outRid
    r.Set("out_rid", _outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRateAddRequest) GetOutRid() string {
    return r._outRid
}
// RateSwitchCal Setter
// 在添加新rate时，同时添加rate开关日历。可以只设定想设定的某些天，可以不连续。date：开关状态控制的是那一天rate_status：开关状态。0，关闭；1，打开
func (r *TaobaoXhotelRateAddRequest) SetRateSwitchCal(_rateSwitchCal string) error {
    r._rateSwitchCal = _rateSwitchCal
    r.Set("rate_switch_cal", _rateSwitchCal)
    return nil
}

// RateSwitchCal Getter
func (r TaobaoXhotelRateAddRequest) GetRateSwitchCal() string {
    return r._rateSwitchCal
}
// LockEndTime Setter
// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoXhotelRateAddRequest) SetLockEndTime(_lockEndTime string) error {
    r._lockEndTime = _lockEndTime
    r.Set("lock_end_time", _lockEndTime)
    return nil
}

// LockEndTime Getter
func (r TaobaoXhotelRateAddRequest) GetLockEndTime() string {
    return r._lockEndTime
}
// LockStartTime Setter
// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoXhotelRateAddRequest) SetLockStartTime(_lockStartTime string) error {
    r._lockStartTime = _lockStartTime
    r.Set("lock_start_time", _lockStartTime)
    return nil
}

// LockStartTime Getter
func (r TaobaoXhotelRateAddRequest) GetLockStartTime() string {
    return r._lockStartTime
}
// CurrencyCodeName Setter
// 币种信息,默认是CNY,  @see com.taobao.trip.hotel.model.enums.CurrencyEnum
func (r *TaobaoXhotelRateAddRequest) SetCurrencyCodeName(_currencyCodeName string) error {
    r._currencyCodeName = _currencyCodeName
    r.Set("currency_code_name", _currencyCodeName)
    return nil
}

// CurrencyCodeName Getter
func (r TaobaoXhotelRateAddRequest) GetCurrencyCodeName() string {
    return r._currencyCodeName
}
// Operator Setter
// 操作人信息
func (r *TaobaoXhotelRateAddRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelRateAddRequest) GetOperator() string {
    return r._operator
}
// Source Setter
// 默认是2 ,
func (r *TaobaoXhotelRateAddRequest) SetSource(_source int64) error {
    r._source = _source
    r.Set("source", _source)
    return nil
}

// Source Getter
func (r TaobaoXhotelRateAddRequest) GetSource() int64 {
    return r._source
}
// Status Setter
// 1是开,0是关, 不填默认是开, rate状态
func (r *TaobaoXhotelRateAddRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelRateAddRequest) GetStatus() int64 {
    return r._status
}
// OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoXhotelRateAddRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
    r._onlineBookingBindingInfo = _onlineBookingBindingInfo
    r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
    return nil
}

// OnlineBookingBindingInfo Getter
func (r TaobaoXhotelRateAddRequest) GetOnlineBookingBindingInfo() string {
    return r._onlineBookingBindingInfo
}
