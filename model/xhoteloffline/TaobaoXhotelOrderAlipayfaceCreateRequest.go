package xhoteloffline

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住支付创建接口 API请求
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

// 初始化TaobaoXhotelOrderAlipayfaceCreateRequest对象
func NewTaobaoXhotelOrderAlipayfaceCreateRequest() *TaobaoXhotelOrderAlipayfaceCreateRequest{
    return &TaobaoXhotelOrderAlipayfaceCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.alipayface.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TotalFee Setter
// 总房费,单位为分
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetTotalFee(totalFee int64) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

// TotalFee Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetTotalFee() int64 {
    return r.totalFee
}
// CheckOut Setter
// 离店日期(最多允许9天)
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetCheckOut(checkOut string) error {
    r.checkOut = checkOut
    r.Set("check_out", checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetCheckOut() string {
    return r.checkOut
}
// HotelCode Setter
// 发布到阿里旅行的酒店编码
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetHotelCode() string {
    return r.hotelCode
}
// CheckIn Setter
// 入住日期
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetCheckIn(checkIn string) error {
    r.checkIn = checkIn
    r.Set("check_in", checkIn)
    return nil
}

// CheckIn Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetCheckIn() string {
    return r.checkIn
}
// DailyPriceInfo Setter
// 每日房价,json格式
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetDailyPriceInfo(dailyPriceInfo string) error {
    r.dailyPriceInfo = dailyPriceInfo
    r.Set("daily_price_info", dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetDailyPriceInfo() string {
    return r.dailyPriceInfo
}
// OutOrderId Setter
// 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetOutOrderId() string {
    return r.outOrderId
}
// RoomQuantity Setter
// 预定的房间数量
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRoomQuantity(roomQuantity int64) error {
    r.roomQuantity = roomQuantity
    r.Set("room_quantity", roomQuantity)
    return nil
}

// RoomQuantity Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRoomQuantity() int64 {
    return r.roomQuantity
}
// Guests Setter
// 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetGuests(guests []Guest) error {
    r.guests = guests
    r.Set("guests", guests)
    return nil
}

// Guests Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetGuests() []Guest {
    return r.guests
}
// AlipayNumber Setter
// 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetAlipayNumber(alipayNumber string) error {
    r.alipayNumber = alipayNumber
    r.Set("alipay_number", alipayNumber)
    return nil
}

// AlipayNumber Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetAlipayNumber() string {
    return r.alipayNumber
}
// Channel Setter
// 订单渠道信息,可以留空
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetChannel() string {
    return r.channel
}
// Vendor Setter
// 不清楚请留空, 用于和outHid共同定位一个酒店
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetVendor() string {
    return r.vendor
}
// RoomtypeName Setter
// 房型名称
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRoomtypeName(roomtypeName string) error {
    r.roomtypeName = roomtypeName
    r.Set("roomtype_name", roomtypeName)
    return nil
}

// RoomtypeName Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRoomtypeName() string {
    return r.roomtypeName
}
// RateplanName Setter
// rateplan名称(不清楚可以留空)
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRateplanName(rateplanName string) error {
    r.rateplanName = rateplanName
    r.Set("rateplan_name", rateplanName)
    return nil
}

// RateplanName Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRateplanName() string {
    return r.rateplanName
}
// SelfCheckin Setter
// 是否为自助入住模式下创建订单，是：true,否：false
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetSelfCheckin(selfCheckin bool) error {
    r.selfCheckin = selfCheckin
    r.Set("self_checkin", selfCheckin)
    return nil
}

// SelfCheckin Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetSelfCheckin() bool {
    return r.selfCheckin
}
