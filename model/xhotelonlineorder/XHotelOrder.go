package xhotelonlineorder

// XHotelOrder 结构体
type XHotelOrder struct {
	// 入住人信息
	Guests []XOrderGuest `json:"guests,omitempty" xml:"guests>x_order_guest,omitempty"`
	// 下单时每间夜的价格（分）
	Prices []int64 `json:"prices,omitempty" xml:"prices>int64,omitempty"`
	// 入住时间
	CheckinDate string `json:"checkin_date,omitempty" xml:"checkin_date,omitempty"`
	// 离店时间
	CheckoutDate string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	// 卖家淘宝账号
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 买家淘宝账号
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 交易状态。 WAIT_BUYER_PAY:预订中/等待买家付款, WAIT_SELLER_SEND_GOODS:预订中/等待卖家发货(确认), TRADE_CLOSED:结束/预订失败/交易关闭, TRADE_FINISHED:结束/交易成功, TRADE_NO_CREATE_PAY:结束/预订失败/没有创建支付宝交易, TRADE_CLOSED_BY_TAOBAO:结束/预订失败/预订被卖家关闭, TRADE_SUCCESS:交易中/预订成功/卖家已确认, TRADE_CHECKIN:交易中/预定成功/买家入住, TRADE_CHECKOUT:交易中/预定成功/买家离店, TRADE_SETTLEING:交易中/预定成功/结账中, TRADE_SETTLE_SUCCESS:结束/预定成功/结账成功
	TradeStatus string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 订单创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 订单修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 买家留言
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 支付宝交易号，28位字符
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 合作方订单号,最长250个字符
	OutOid string `json:"out_oid,omitempty" xml:"out_oid,omitempty"`
	// 买家最早到达时间
	ArriveEarly string `json:"arrive_early,omitempty" xml:"arrive_early,omitempty"`
	// 买家最晚到达时间
	ArriveLate string `json:"arrive_late,omitempty" xml:"arrive_late,omitempty"`
	// logisticsStatus
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// 预付订单使用，1买家已经申请退款，等待卖家同意，2卖家已经同意退款，等待买家退货，3买家已经退货，等待卖家确认收货，4退款关闭，5退款成功，6卖家拒绝退款，7没有申请退款
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 酒店订单id
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 酒店id
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 房型id
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// RatePlan 中的 rpid
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
	// tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 天数
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 支付类型 可选值 1：预付 5：前台面付
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 总房价（分）
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 实付款（分）
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
}
