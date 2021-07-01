package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格推送接口（全量更新） API请求
taobao.xhotel.rate.update

酒店产品库rate更新
*/
type TaobaoXhotelRateUpdateAPIRequest struct {
    model.Params
    // 每日价格和房价专有库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
    _inventoryPrice   string
    // 商家价格计划编码
    _rateplanCode   string
    // 商家房型ID
    _outRid   string
    // 系统商，一般不用填写，使用需要申请
    _vendor   string
    // 日历价格开关， date：开关状态控制的是那一天 rate_status：开关状态。0，关闭；1，打开
    _rateSwitchCal   string
    // 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    _lockEndTime   string
    // 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    _lockStartTime   string
    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    _onlineBookingBindingInfo   string
}

// 初始化TaobaoXhotelRateUpdateAPIRequest对象
func NewTaobaoXhotelRateUpdateRequest() *TaobaoXhotelRateUpdateAPIRequest{
    return &TaobaoXhotelRateUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InventoryPrice Setter
// 每日价格和房价专有库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
func (r *TaobaoXhotelRateUpdateAPIRequest) SetInventoryPrice(_inventoryPrice string) error {
    r._inventoryPrice = _inventoryPrice
    r.Set("inventory_price", _inventoryPrice)
    return nil
}

// InventoryPrice Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetInventoryPrice() string {
    return r._inventoryPrice
}
// RateplanCode Setter
// 商家价格计划编码
func (r *TaobaoXhotelRateUpdateAPIRequest) SetRateplanCode(_rateplanCode string) error {
    r._rateplanCode = _rateplanCode
    r.Set("rateplan_code", _rateplanCode)
    return nil
}

// RateplanCode Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetRateplanCode() string {
    return r._rateplanCode
}
// OutRid Setter
// 商家房型ID
func (r *TaobaoXhotelRateUpdateAPIRequest) SetOutRid(_outRid string) error {
    r._outRid = _outRid
    r.Set("out_rid", _outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetOutRid() string {
    return r._outRid
}
// Vendor Setter
// 系统商，一般不用填写，使用需要申请
func (r *TaobaoXhotelRateUpdateAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetVendor() string {
    return r._vendor
}
// RateSwitchCal Setter
// 日历价格开关， date：开关状态控制的是那一天 rate_status：开关状态。0，关闭；1，打开
func (r *TaobaoXhotelRateUpdateAPIRequest) SetRateSwitchCal(_rateSwitchCal string) error {
    r._rateSwitchCal = _rateSwitchCal
    r.Set("rate_switch_cal", _rateSwitchCal)
    return nil
}

// RateSwitchCal Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetRateSwitchCal() string {
    return r._rateSwitchCal
}
// LockEndTime Setter
// 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoXhotelRateUpdateAPIRequest) SetLockEndTime(_lockEndTime string) error {
    r._lockEndTime = _lockEndTime
    r.Set("lock_end_time", _lockEndTime)
    return nil
}

// LockEndTime Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetLockEndTime() string {
    return r._lockEndTime
}
// LockStartTime Setter
// 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
func (r *TaobaoXhotelRateUpdateAPIRequest) SetLockStartTime(_lockStartTime string) error {
    r._lockStartTime = _lockStartTime
    r.Set("lock_start_time", _lockStartTime)
    return nil
}

// LockStartTime Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetLockStartTime() string {
    return r._lockStartTime
}
// OnlineBookingBindingInfo Setter
// 在线预约关联关系推送，priceRuleNumber：加价规则序号
func (r *TaobaoXhotelRateUpdateAPIRequest) SetOnlineBookingBindingInfo(_onlineBookingBindingInfo string) error {
    r._onlineBookingBindingInfo = _onlineBookingBindingInfo
    r.Set("online_booking_binding_info", _onlineBookingBindingInfo)
    return nil
}

// OnlineBookingBindingInfo Getter
func (r TaobaoXhotelRateUpdateAPIRequest) GetOnlineBookingBindingInfo() string {
    return r._onlineBookingBindingInfo
}
