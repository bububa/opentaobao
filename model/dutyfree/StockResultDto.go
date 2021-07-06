package dutyfree

// StockResultDto 结构体
type StockResultDto struct {
	// 条形码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
}
