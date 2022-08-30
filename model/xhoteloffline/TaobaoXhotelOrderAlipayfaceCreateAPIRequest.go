package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceCreateAPIRequest 信用住支付创建接口 API请求
// taobao.xhotel.order.alipayface.create
//
// 用于创建一笔信用住支付，主要应用场景是线下信用住
type TaobaoXhotelOrderAlipayfaceCreateAPIRequest struct {
	model.Params
	// 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
	_guests []Guest
	// 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
	_outOrderId string
	// 发布到阿里旅行的酒店编码
	_hotelCode string
	// 入住日期
	_checkIn string
	// 离店日期(最多允许9天)
	_checkOut string
	// 每日房价,json格式
	_dailyPriceInfo string
	// 房型名称
	_roomtypeName string
	// rateplan名称(不清楚可以留空)
	_rateplanName string
	// 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
	_alipayNumber string
	// 订单渠道信息,可以留空
	_channel string
	// 不清楚请留空, 用于和outHid共同定位一个酒店
	_vendor string
	// 预定的房间数量
	_roomQuantity int64
	// 总房费,单位为分
	_totalFee int64
	// 是否为自助入住模式下创建订单，是：true,否：false
	_selfCheckin bool
}

// NewTaobaoXhotelOrderAlipayfaceCreateRequest 初始化TaobaoXhotelOrderAlipayfaceCreateAPIRequest对象
func NewTaobaoXhotelOrderAlipayfaceCreateRequest() *TaobaoXhotelOrderAlipayfaceCreateAPIRequest {
	return &TaobaoXhotelOrderAlipayfaceCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.alipayface.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGuests is Guests Setter
// 入住人信息, 注意必须有且只有一个设置为主入住人, 用于信用住结算扣款. 对于java版本的SDK可以使用setGuests(List &lt Guest &gt guests)赋值; 对于.net等其他版本SDK可以通过将List &lt Guest&gt结构数据转为json串赋值.
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetGuests(_guests []Guest) error {
	r._guests = _guests
	r.Set("guests", _guests)
	return nil
}

// GetGuests Guests Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetGuests() []Guest {
	return r._guests
}

// SetOutOrderId is OutOrderId Setter
// 商家系统的订单号，必须全局唯一，重复会按照相同订单处理
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetHotelCode is HotelCode Setter
// 发布到阿里旅行的酒店编码
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期(最多允许9天)
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetDailyPriceInfo is DailyPriceInfo Setter
// 每日房价,json格式
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetDailyPriceInfo(_dailyPriceInfo string) error {
	r._dailyPriceInfo = _dailyPriceInfo
	r.Set("daily_price_info", _dailyPriceInfo)
	return nil
}

// GetDailyPriceInfo DailyPriceInfo Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetDailyPriceInfo() string {
	return r._dailyPriceInfo
}

// SetRoomtypeName is RoomtypeName Setter
// 房型名称
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetRoomtypeName(_roomtypeName string) error {
	r._roomtypeName = _roomtypeName
	r.Set("roomtype_name", _roomtypeName)
	return nil
}

// GetRoomtypeName RoomtypeName Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetRoomtypeName() string {
	return r._roomtypeName
}

// SetRateplanName is RateplanName Setter
// rateplan名称(不清楚可以留空)
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetRateplanName(_rateplanName string) error {
	r._rateplanName = _rateplanName
	r.Set("rateplan_name", _rateplanName)
	return nil
}

// GetRateplanName RateplanName Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetRateplanName() string {
	return r._rateplanName
}

// SetAlipayNumber is AlipayNumber Setter
// 扫描用户支付宝得到的串号, 该字段不为空时会采用串号对应的支付宝账号进行信用住结算
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetAlipayNumber(_alipayNumber string) error {
	r._alipayNumber = _alipayNumber
	r.Set("alipay_number", _alipayNumber)
	return nil
}

// GetAlipayNumber AlipayNumber Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetAlipayNumber() string {
	return r._alipayNumber
}

// SetChannel is Channel Setter
// 订单渠道信息,可以留空
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetChannel(_channel string) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetChannel() string {
	return r._channel
}

// SetVendor is Vendor Setter
// 不清楚请留空, 用于和outHid共同定位一个酒店
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRoomQuantity is RoomQuantity Setter
// 预定的房间数量
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetRoomQuantity(_roomQuantity int64) error {
	r._roomQuantity = _roomQuantity
	r.Set("room_quantity", _roomQuantity)
	return nil
}

// GetRoomQuantity RoomQuantity Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetRoomQuantity() int64 {
	return r._roomQuantity
}

// SetTotalFee is TotalFee Setter
// 总房费,单位为分
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetTotalFee(_totalFee int64) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetTotalFee() int64 {
	return r._totalFee
}

// SetSelfCheckin is SelfCheckin Setter
// 是否为自助入住模式下创建订单，是：true,否：false
func (r *TaobaoXhotelOrderAlipayfaceCreateAPIRequest) SetSelfCheckin(_selfCheckin bool) error {
	r._selfCheckin = _selfCheckin
	r.Set("self_checkin", _selfCheckin)
	return nil
}

// GetSelfCheckin SelfCheckin Getter
func (r TaobaoXhotelOrderAlipayfaceCreateAPIRequest) GetSelfCheckin() bool {
	return r._selfCheckin
}
