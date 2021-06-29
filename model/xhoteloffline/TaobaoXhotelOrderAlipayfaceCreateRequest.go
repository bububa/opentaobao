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
    _totalFee   int64
    // 离店日期(最多允许9天)
    _checkOut   string
    // 发布到阿里旅行的酒店编码
    _hotelCode   string
    // 入住日期
    _checkIn   string
    // 每日房价,json格式
    _dailyPriceInfo   string
    // 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
    _outOrderId   string
    // 预定的房间数量
    _roomQuantity   int64
    // 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
    _guests   []Guest
    // 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
    _alipayNumber   string
    // 订单渠道信息,可以留空
    _channel   string
    // 不清楚请留空, 用于和outHid共同定位一个酒店
    _vendor   string
    // 房型名称
    _roomtypeName   string
    // rateplan名称(不清楚可以留空)
    _rateplanName   string
    // 是否为自助入住模式下创建订单，是：true,否：false
    _selfCheckin   bool
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
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetTotalFee(_totalFee int64) error {
    r._totalFee = _totalFee
    r.Set("total_fee", _totalFee)
    return nil
}

// TotalFee Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetTotalFee() int64 {
    return r._totalFee
}
// CheckOut Setter
// 离店日期(最多允许9天)
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetCheckOut(_checkOut string) error {
    r._checkOut = _checkOut
    r.Set("check_out", _checkOut)
    return nil
}

// CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetCheckOut() string {
    return r._checkOut
}
// HotelCode Setter
// 发布到阿里旅行的酒店编码
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetHotelCode() string {
    return r._hotelCode
}
// CheckIn Setter
// 入住日期
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetCheckIn(_checkIn string) error {
    r._checkIn = _checkIn
    r.Set("check_in", _checkIn)
    return nil
}

// CheckIn Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetCheckIn() string {
    return r._checkIn
}
// DailyPriceInfo Setter
// 每日房价,json格式
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
    r._dailyPriceInfo = _dailyPriceInfo
    r.Set("daily_price_info", _dailyPriceInfo)
    return nil
}

// DailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetDailyPriceInfo() string {
    return r._dailyPriceInfo
}
// OutOrderId Setter
// 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetOutOrderId() string {
    return r._outOrderId
}
// RoomQuantity Setter
// 预定的房间数量
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRoomQuantity(_roomQuantity int64) error {
    r._roomQuantity = _roomQuantity
    r.Set("room_quantity", _roomQuantity)
    return nil
}

// RoomQuantity Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRoomQuantity() int64 {
    return r._roomQuantity
}
// Guests Setter
// 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetGuests(_guests []Guest) error {
    r._guests = _guests
    r.Set("guests", _guests)
    return nil
}

// Guests Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetGuests() []Guest {
    return r._guests
}
// AlipayNumber Setter
// 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetAlipayNumber(_alipayNumber string) error {
    r._alipayNumber = _alipayNumber
    r.Set("alipay_number", _alipayNumber)
    return nil
}

// AlipayNumber Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetAlipayNumber() string {
    return r._alipayNumber
}
// Channel Setter
// 订单渠道信息,可以留空
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetChannel() string {
    return r._channel
}
// Vendor Setter
// 不清楚请留空, 用于和outHid共同定位一个酒店
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetVendor() string {
    return r._vendor
}
// RoomtypeName Setter
// 房型名称
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRoomtypeName(_roomtypeName string) error {
    r._roomtypeName = _roomtypeName
    r.Set("roomtype_name", _roomtypeName)
    return nil
}

// RoomtypeName Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRoomtypeName() string {
    return r._roomtypeName
}
// RateplanName Setter
// rateplan名称(不清楚可以留空)
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetRateplanName(_rateplanName string) error {
    r._rateplanName = _rateplanName
    r.Set("rateplan_name", _rateplanName)
    return nil
}

// RateplanName Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetRateplanName() string {
    return r._rateplanName
}
// SelfCheckin Setter
// 是否为自助入住模式下创建订单，是：true,否：false
func (r *TaobaoXhotelOrderAlipayfaceCreateRequest) SetSelfCheckin(_selfCheckin bool) error {
    r._selfCheckin = _selfCheckin
    r.Set("self_checkin", _selfCheckin)
    return nil
}

// SelfCheckin Getter
func (r TaobaoXhotelOrderAlipayfaceCreateRequest) GetSelfCheckin() bool {
    return r._selfCheckin
}
