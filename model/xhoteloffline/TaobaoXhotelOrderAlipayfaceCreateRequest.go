package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住支付创建接口 APIRequest
taobao.xhotel.order.alipayface.create

用于创建一笔信用住支付，主要应用场景是线下信用住
*/
type TaobaoXhotelOrderAlipayfaceCreateRequest struct {
    model.Params

    // 总房费,单位为分
    totalFee   int64 

    // 离店日期(最多允许9天)
    checkOut   string 

    // 发布到阿里旅行的酒店编码
    hotelCode   string 

    // 入住日期
    checkIn   string 

    // 每日房价,json格式
    dailyPriceInfo   string 

    // 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
    outOrderId   string 

    // 预定的房间数量
    roomQuantity   int64 

    // 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
    guests   []Guest 

    // 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
    alipayNumber   string 

    // 订单渠道信息,可以留空
    channel   string 

    // 不清楚请留空, 用于和outHid共同定位一个酒店
    vendor   string 

    // 房型名称
    roomtypeName   string 

    // rateplan名称(不清楚可以留空)
    rateplanName   string 

    // 是否为自助入住模式下创建订单，是：true,否：false
    selfCheckin   bool 

}

func NewTaobaoXhotelOrderAlipayfaceCreateRequest() *TaobaoXhotelOrderAlipayfaceCreateRequest{
    return &TaobaoXhotelOrderAlipayfaceCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.create"
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetTotalFee() int64 {
    return r.totalFee
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetCheckOut() string {
    return r.checkOut
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetCheckIn(checkIn string) error {
    r.checkIn = checkIn
    r.Set("check_in", checkIn)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetCheckIn() string {
    return r.checkIn
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetDailyPriceInfo(dailyPriceInfo string) error {
    r.dailyPriceInfo = dailyPriceInfo
    r.Set("daily_price_info", dailyPriceInfo)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetDailyPriceInfo() string {
    return r.dailyPriceInfo
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetOutOrderId() string {
    return r.outOrderId
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRoomQuantity(roomQuantity int64) error {
    r.roomQuantity = roomQuantity
    r.Set("room_quantity", roomQuantity)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRoomQuantity() int64 {
    return r.roomQuantity
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetGuests(guests []Guest) error {
    r.guests = guests
    r.Set("guests", guests)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetGuests() []Guest {
    return r.guests
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetAlipayNumber(alipayNumber string) error {
    r.alipayNumber = alipayNumber
    r.Set("alipay_number", alipayNumber)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetAlipayNumber() string {
    return r.alipayNumber
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetChannel() string {
    return r.channel
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRoomtypeName(roomtypeName string) error {
    r.roomtypeName = roomtypeName
    r.Set("roomtype_name", roomtypeName)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRoomtypeName() string {
    return r.roomtypeName
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRateplanName(rateplanName string) error {
    r.rateplanName = rateplanName
    r.Set("rateplan_name", rateplanName)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRateplanName() string {
    return r.rateplanName
}

func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetSelfCheckin(selfCheckin bool) error {
    r.selfCheckin = selfCheckin
    r.Set("self_checkin", selfCheckin)
    return nil
}

func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetSelfCheckin() bool {
    return r.selfCheckin
}

