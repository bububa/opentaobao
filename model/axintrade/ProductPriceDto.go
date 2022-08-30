package axintrade

// ProductPriceDto 结构体
type ProductPriceDto struct {
	// 日期 格式yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 价格，单位为分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
