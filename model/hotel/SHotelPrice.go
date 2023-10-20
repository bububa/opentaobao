package hotel

// ShotelPrice 结构体
type ShotelPrice struct {
	// 每个标准酒店某一天的所有库价集合
	DailyPriceList []ShotelDailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>shotel_daily_price,omitempty"`
	// 标准房型shid
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}
