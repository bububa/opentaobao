package bus

// TopItemChangeNotifyRq 结构体
type TopItemChangeNotifyRq struct {
	// 出发城市
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// 目的地城市
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
	// 发车日期，格式：yyyy-MM-dd
	FromDate string `json:"from_date,omitempty" xml:"from_date,omitempty"`
}
