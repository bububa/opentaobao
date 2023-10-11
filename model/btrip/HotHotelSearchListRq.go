package btrip

// HotHotelSearchListRq 结构体
type HotHotelSearchListRq struct {
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 分页大小,不能超过50（默认20）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 首次传空，之后每次请求传入上次返回的hotelId。 这个和分页只能用一种如果俩个都传递就按照hotelId来查询
	LastHotelId int64 `json:"last_hotel_id,omitempty" xml:"last_hotel_id,omitempty"`
}
