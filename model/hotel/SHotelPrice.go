package hotel

// SHotelPrice 结构体
type SHotelPrice struct {
	// 每个标准酒店某一天的所有库价集合
	DailyPriceList []SHotelDailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>s_hotel_daily_price,omitempty"`
	// 标准房型shid
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
