package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
价格推送接口（全量更新） APIRequest
taobao.xhotel.rate.update

酒店产品库rate更新
*/
type TaobaoXhotelRateUpdateRequest struct {
    model.Params

    // 每日价格和房价专有库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+180 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
    inventoryPrice   string 

    // 商家价格计划编码
    rateplanCode   string 

    // 商家房型ID
    outRid   string 

    // 系统商，一般不用填写，使用需要申请
    vendor   string 

    // 日历价格开关， date：开关状态控制的是那一天 rate_status：开关状态。0，关闭；1，打开
    rateSwitchCal   string 

    // 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    lockEndTime   string 

    // 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    lockStartTime   string 

    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    onlineBookingBindingInfo   string 

}

func NewTaobaoXhotelRateUpdateRequest() *TaobaoXhotelRateUpdateRequest{
    return &TaobaoXhotelRateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.update"
}

func (r TaobaoXhotelRateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateUpdateRequest) SetInventoryPrice(inventoryPrice string) error {
    r.inventoryPrice = inventoryPrice
    r.Set("inventory_price", inventoryPrice)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetInventoryPrice() string {
    return r.inventoryPrice
}

func (r *TaobaoXhotelRateUpdateRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetRateplanCode() string {
    return r.rateplanCode
}

func (r *TaobaoXhotelRateUpdateRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelRateUpdateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRateUpdateRequest) SetRateSwitchCal(rateSwitchCal string) error {
    r.rateSwitchCal = rateSwitchCal
    r.Set("rate_switch_cal", rateSwitchCal)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetRateSwitchCal() string {
    return r.rateSwitchCal
}

func (r *TaobaoXhotelRateUpdateRequest) SetLockEndTime(lockEndTime string) error {
    r.lockEndTime = lockEndTime
    r.Set("lock_end_time", lockEndTime)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetLockEndTime() string {
    return r.lockEndTime
}

func (r *TaobaoXhotelRateUpdateRequest) SetLockStartTime(lockStartTime string) error {
    r.lockStartTime = lockStartTime
    r.Set("lock_start_time", lockStartTime)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetLockStartTime() string {
    return r.lockStartTime
}

func (r *TaobaoXhotelRateUpdateRequest) SetOnlineBookingBindingInfo(onlineBookingBindingInfo string) error {
    r.onlineBookingBindingInfo = onlineBookingBindingInfo
    r.Set("online_booking_binding_info", onlineBookingBindingInfo)
    return nil
}

func (r TaobaoXhotelRateUpdateRequest) GetOnlineBookingBindingInfo() string {
    return r.onlineBookingBindingInfo
}

