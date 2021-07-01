package tuanhotel

// TopSkuDailyInfo 结构体
type TopSkuDailyInfo struct {
	// 日期
	D string `json:"d,omitempty" xml:"d,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}
