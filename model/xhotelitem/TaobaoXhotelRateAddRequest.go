package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增专享房价 APIRequest
taobao.xhotel.rate.add

酒店产品库rate添加
*/
type TaobaoXhotelRateAddRequest struct {
    model.Params

    // gid酒店商品id
    gid   int64 

    // 酒店RPID
    rpid   int64 

    // 名称
    name   string 

    // 额外服务-是否可以加床，1：不可以，2：可以
    addBed   int64 

    // 额外服务-加床价格
    addBedPrice   int64 

    // 币种（仅支持CNY）
    currencyCode   int64 

    // 实价有房标签（RP支付类型为全额支付）
    shijiaTag   int64 

    // 价格和库存信息。A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复C:price 价格 int类型 取值范围1-99999999 单位为分D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
    inventoryPrice   string 

    // “即时确认”标识，此类商品预订后直接发货。
    jishiquerenTag   int64 

    // 卖家自己系统的Code，简称RateCode
    rateplanCode   string 

    // 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
    vendor   string 

    // 卖家房型ID, 这是卖家自己系统中的房型ID，注意：需按照规则组合
    outRid   string 

    // 在添加新rate时，同时添加rate开关日历。可以只设定想设定的某些天，可以不连续。date：开关状态控制的是那一天rate_status：开关状态。0，关闭；1，打开
    rateSwitchCal   string 

    // 锁库存截止时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    lockEndTime   string 

    // 锁库存开始时间，如果当前时间是在锁库存开始时间和截止时间之间，那么不允许修改该活动库存（包含开始时间和截止时间）
    lockStartTime   string 

    // 币种信息,默认是CNY,  @see com.taobao.trip.hotel.model.enums.CurrencyEnum
    currencyCodeName   string 

    // 操作人信息
    operator   string 

    // 默认是2 ,
    source   int64 

    // 1是开,0是关, 不填默认是开, rate状态
    status   int64 

    // 在线预约关联关系推送，priceRuleNumber：加价规则序号
    onlineBookingBindingInfo   string 

}

func NewTaobaoXhotelRateAddRequest() *TaobaoXhotelRateAddRequest{
    return &TaobaoXhotelRateAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRateAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.rate.add"
}

func (r TaobaoXhotelRateAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRateAddRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetGid() int64 {
    return r.gid
}

func (r *TaobaoXhotelRateAddRequest) SetRpid(rpid int64) error {
    r.rpid = rpid
    r.Set("rpid", rpid)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetRpid() int64 {
    return r.rpid
}

func (r *TaobaoXhotelRateAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoXhotelRateAddRequest) SetAddBed(addBed int64) error {
    r.addBed = addBed
    r.Set("add_bed", addBed)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetAddBed() int64 {
    return r.addBed
}

func (r *TaobaoXhotelRateAddRequest) SetAddBedPrice(addBedPrice int64) error {
    r.addBedPrice = addBedPrice
    r.Set("add_bed_price", addBedPrice)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetAddBedPrice() int64 {
    return r.addBedPrice
}

func (r *TaobaoXhotelRateAddRequest) SetCurrencyCode(currencyCode int64) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetCurrencyCode() int64 {
    return r.currencyCode
}

func (r *TaobaoXhotelRateAddRequest) SetShijiaTag(shijiaTag int64) error {
    r.shijiaTag = shijiaTag
    r.Set("shijia_tag", shijiaTag)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetShijiaTag() int64 {
    return r.shijiaTag
}

func (r *TaobaoXhotelRateAddRequest) SetInventoryPrice(inventoryPrice string) error {
    r.inventoryPrice = inventoryPrice
    r.Set("inventory_price", inventoryPrice)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetInventoryPrice() string {
    return r.inventoryPrice
}

func (r *TaobaoXhotelRateAddRequest) SetJishiquerenTag(jishiquerenTag int64) error {
    r.jishiquerenTag = jishiquerenTag
    r.Set("jishiqueren_tag", jishiquerenTag)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetJishiquerenTag() int64 {
    return r.jishiquerenTag
}

func (r *TaobaoXhotelRateAddRequest) SetRateplanCode(rateplanCode string) error {
    r.rateplanCode = rateplanCode
    r.Set("rateplan_code", rateplanCode)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetRateplanCode() string {
    return r.rateplanCode
}

func (r *TaobaoXhotelRateAddRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelRateAddRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetOutRid() string {
    return r.outRid
}

func (r *TaobaoXhotelRateAddRequest) SetRateSwitchCal(rateSwitchCal string) error {
    r.rateSwitchCal = rateSwitchCal
    r.Set("rate_switch_cal", rateSwitchCal)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetRateSwitchCal() string {
    return r.rateSwitchCal
}

func (r *TaobaoXhotelRateAddRequest) SetLockEndTime(lockEndTime string) error {
    r.lockEndTime = lockEndTime
    r.Set("lock_end_time", lockEndTime)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetLockEndTime() string {
    return r.lockEndTime
}

func (r *TaobaoXhotelRateAddRequest) SetLockStartTime(lockStartTime string) error {
    r.lockStartTime = lockStartTime
    r.Set("lock_start_time", lockStartTime)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetLockStartTime() string {
    return r.lockStartTime
}

func (r *TaobaoXhotelRateAddRequest) SetCurrencyCodeName(currencyCodeName string) error {
    r.currencyCodeName = currencyCodeName
    r.Set("currency_code_name", currencyCodeName)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetCurrencyCodeName() string {
    return r.currencyCodeName
}

func (r *TaobaoXhotelRateAddRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetOperator() string {
    return r.operator
}

func (r *TaobaoXhotelRateAddRequest) SetSource(source int64) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetSource() int64 {
    return r.source
}

func (r *TaobaoXhotelRateAddRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetStatus() int64 {
    return r.status
}

func (r *TaobaoXhotelRateAddRequest) SetOnlineBookingBindingInfo(onlineBookingBindingInfo string) error {
    r.onlineBookingBindingInfo = onlineBookingBindingInfo
    r.Set("online_booking_binding_info", onlineBookingBindingInfo)
    return nil
}

func (r TaobaoXhotelRateAddRequest) GetOnlineBookingBindingInfo() string {
    return r.onlineBookingBindingInfo
}

