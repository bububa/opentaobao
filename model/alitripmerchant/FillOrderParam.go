package alitripmerchant

// FillOrderParam 结构体
type FillOrderParam struct {
	// 离店时间
	CheckOutDate string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
	// 外部酒店id
	HotelId string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// 商品名称
	RpTitle string `json:"rp_title,omitempty" xml:"rp_title,omitempty"`
	// 入住时间
	CheckInDate string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// 代金券id
	VoucherId string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
	// 标准静态信息库酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 房型id
	SrId int64 `json:"sr_id,omitempty" xml:"sr_id,omitempty"`
}
