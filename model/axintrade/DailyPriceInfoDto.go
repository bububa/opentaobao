package axintrade

// DailyPriceInfoDto 结构体
type DailyPriceInfoDto struct {
	// 日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 餐食
	Board *Board `json:"board,omitempty" xml:"board,omitempty"`
	// 房间价格（人民币元）
	CnyPrice int64 `json:"cny_price,omitempty" xml:"cny_price,omitempty"`
	// 原币种金额
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
