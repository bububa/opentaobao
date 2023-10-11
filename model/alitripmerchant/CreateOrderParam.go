package alitripmerchant

// CreateOrderParam 结构体
type CreateOrderParam struct {
	// 多房间参数
	GuestByRoomDtos []GuestByRoomDto `json:"guest_by_room_dtos,omitempty" xml:"guest_by_room_dtos>guest_by_room_dto,omitempty"`
	// 早餐名称
	BreakfastName string `json:"breakfast_name,omitempty" xml:"breakfast_name,omitempty"`
	// 会员可见等级
	MemberLevel string `json:"member_level,omitempty" xml:"member_level,omitempty"`
	// 外部酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 支付渠道
	PaymentChannel string `json:"payment_channel,omitempty" xml:"payment_channel,omitempty"`
	// 入住人电话
	CustomerPhone string `json:"customer_phone,omitempty" xml:"customer_phone,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住人名
	CustomerFirstName string `json:"customer_first_name,omitempty" xml:"customer_first_name,omitempty"`
	// 入住人邮箱
	CustomerEmail string `json:"customer_email,omitempty" xml:"customer_email,omitempty"`
	// 入住人姓
	CustomerLastName string `json:"customer_last_name,omitempty" xml:"customer_last_name,omitempty"`
	// 价格名称
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 活动标识字段
	OfferSourceChannel string `json:"offer_source_channel,omitempty" xml:"offer_source_channel,omitempty"`
	// 活动id
	OfferId string `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// rpcode
	RpCode string `json:"rp_code,omitempty" xml:"rp_code,omitempty"`
	// 代金券id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 1
	AreaCode string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	// 外部outRoomId
	OutRoomId int64 `json:"out_room_id,omitempty" xml:"out_room_id,omitempty"`
	// 宝贝ID
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 预订酒店房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 标准静态信息库酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 儿童数量
	ChildrenNumber int64 `json:"children_number,omitempty" xml:"children_number,omitempty"`
	// 内部rpId
	RpId int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// 成人数量
	PersonNumber int64 `json:"person_number,omitempty" xml:"person_number,omitempty"`
	// 成人数量
	AdultNumber int64 `json:"adult_number,omitempty" xml:"adult_number,omitempty"`
	// 房间Id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 价格id
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 支付类型
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 优惠券实例id
	CouponInstanceId int64 `json:"coupon_instance_id,omitempty" xml:"coupon_instance_id,omitempty"`
	// 权益券相关参数
	DerbyVoucherUniversalDTO *DerbyVoucherUniversalDto `json:"derby_voucher_universal_d_t_o,omitempty" xml:"derby_voucher_universal_d_t_o,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 是否同意向境外提供个人信息
	DataExportAgreement bool `json:"data_export_agreement,omitempty" xml:"data_export_agreement,omitempty"`
}
