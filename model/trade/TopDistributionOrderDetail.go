package trade

// TopDistributionOrderDetail 结构体
type TopDistributionOrderDetail struct {
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单状态
	OrderStatusDesc string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// 房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// 离店日期
	CheckoutDate string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	// rp名称
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 订单更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 订单类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 入住时间
	CheckinDate string `json:"checkin_date,omitempty" xml:"checkin_date,omitempty"`
	// 买家姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 订单创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 分销渠道外部订单号
	DisOid string `json:"dis_oid,omitempty" xml:"dis_oid,omitempty"`
	// 飞猪拆单订单（当一次下单定多间房的时候会出现这种情况）
	MultiTids string `json:"multi_tids,omitempty" xml:"multi_tids,omitempty"`
	// 房型id
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 订单类型id
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 付款金额
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 总房费
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// 订单状态Code
	OrderStatusCode int64 `json:"order_status_code,omitempty" xml:"order_status_code,omitempty"`
	// 酒店hid
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// 飞猪订单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 价格计划ID
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
	// 间夜
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
