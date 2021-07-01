package alitripmerchant

// OrderDto 结构体
type OrderDto struct {
	// 床型名称
	BedName string `json:"bed_name,omitempty" xml:"bed_name,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 支付剩余时间
	PayRemainTime int64 `json:"pay_remain_time,omitempty" xml:"pay_remain_time,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 用户支付总价
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 预订日期
	BookDate string `json:"book_date,omitempty" xml:"book_date,omitempty"`
	// 入住人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 房间数量
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 房型照片
	RoomPhotoUrl string `json:"room_photo_url,omitempty" xml:"room_photo_url,omitempty"`
	// 酒店房型id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 酒店外部id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 成人数量
	AdultNumber int64 `json:"adult_number,omitempty" xml:"adult_number,omitempty"`
	// 儿童数量
	ChildrenNumber int64 `json:"children_number,omitempty" xml:"children_number,omitempty"`
}
