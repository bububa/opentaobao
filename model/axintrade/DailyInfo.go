package axintrade

// DailyInfo 结构体
type DailyInfo struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 价格，单位为分
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
	// 餐食
	BoardDTO *BoardDto `json:"board_d_t_o,omitempty" xml:"board_d_t_o,omitempty"`
	// 原币种价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
