package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
复杂价格推送接口（全量更新） APIRequest
taobao.xhotel.multiplerate.update

酒店产品库复杂rate（多人价，连住价等）更新
同时完全涵盖taobao.xhotel.rate.update的功能
*/
type TaobaoXhotelMultiplerateUpdateRequest struct {
    model.Params

    // 废弃，使用out_rid
    gid   int64 

    // 卖家房型ID
    outRid   string 

    // 废弃，使用rate_plan_code
    rpid   int64 

    // 卖家自己系统的房价code
    ratePlanCode   string 

    // 废弃
    name   string 

    // 入住人数(范围1~10)
    occupancy   int64 

    // 连住天数(范围1~5)
    lengthofstay   int64 

    // 价格和库存信息。 A:use_room_inventory:是否使用房型共享库存，可选值 true false ,false时：使用房价专有库存，此时要求价格和库存必填。 B:date 日期必须为 T---T+180 日内的日期（T为当天），且不能重复 C:price 价格 int类型 取值范围1-99999999 单位为分 D:quota 库存 int 类型 取值范围 0-999（数量库存） 60000(状态库存关) 61000(状态库存开) tax为税费，addBed为加床价，addPerson为加人价1,若连住大于1，price请推送总价
    inventoryPrice   string 

    // 价格开关 date：开关状态控制的那一天；rate_status：开关状态(0，关闭；1，打开); checkin_status：入住开关(0，关闭；1，打开)；checkout_status：离店开关 (0，关闭；1，打开)
    rateSwitchCal   string 

    // 币种.CNY为人民币
    currencyCode   string 

    // 价格状态。0为不可售；1为可售，默认可售
    status   int64 

    // 系统商，一般不填写，使用须申请
    vendor   string 

    // 儿童人数
    childnum   int64 

    // 婴儿人数
    infantnum   int64 

    // 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    lockEndTime   string 

    // 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    lockStartTime   string 

    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    onlineBookingBindingInfo   string 

}

func NewTaobaoXhotelMultiplerateUpdateRequest() *TaobaoXhotelMultiplerateUpdateRequest{
    return &TaobaoXhotelMultiplerateUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.multiplerate.update"
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelMultiplerateUpdateRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetGid() int64 {
    return r.gid
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetRpid() int64 {
    return r.rpid
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetRatePlanCode(ratePlanCode string) error {
    r.ratePlanCode = ratePlanCode
    r.Set("rate_plan_code", ratePlanCode)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetRatePlanCode() string {
    return r.ratePlanCode
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetName() string {
    return r.name
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetOccupancy(occupancy int64) error {
    r.occupancy = occupancy
    r.Set("occupancy", occupancy)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetOccupancy() int64 {
    return r.occupancy
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetLengthofstay(lengthofstay int64) error {
    r.lengthofstay = lengthofstay
    r.Set("lengthofstay", lengthofstay)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetLengthofstay() int64 {
    return r.lengthofstay
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetInventoryPrice(inventoryPrice string) error {
    r.inventoryPrice = inventoryPrice
    r.Set("inventory_price", inventoryPrice)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetInventoryPrice() string {
    return r.inventoryPrice
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetRateSwitchCal(rateSwitchCal string) error {
    r.rateSwitchCal = rateSwitchCal
    r.Set("rate_switch_cal", rateSwitchCal)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetRateSwitchCal() string {
    return r.rateSwitchCal
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetCurrencyCode(currencyCode string) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetCurrencyCode() string {
    return r.currencyCode
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetChildnum(childnum int64) error {
    r.childnum = childnum
    r.Set("childnum", childnum)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetChildnum() int64 {
    return r.childnum
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetInfantnum(infantnum int64) error {
    r.infantnum = infantnum
    r.Set("infantnum", infantnum)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetInfantnum() int64 {
    return r.infantnum
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetLockEndTime(lockEndTime string) error {
    r.lockEndTime = lockEndTime
    r.Set("lock_end_time", lockEndTime)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetLockEndTime() string {
    return r.lockEndTime
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetLockStartTime(lockStartTime string) error {
    r.lockStartTime = lockStartTime
    r.Set("lock_start_time", lockStartTime)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetLockStartTime() string {
    return r.lockStartTime
}

func (r *TaobaoXhotelMultiplerateUpdateRequest) SetOnlineBookingBindingInfo(onlineBookingBindingInfo string) error {
    r.onlineBookingBindingInfo = onlineBookingBindingInfo
    r.Set("online_booking_binding_info", onlineBookingBindingInfo)
    return nil
}

func (r TaobaoXhotelMultiplerateUpdateRequest) GetOnlineBookingBindingInfo() string {
    return r.onlineBookingBindingInfo
}

