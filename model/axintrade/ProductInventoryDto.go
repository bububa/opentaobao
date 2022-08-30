package axintrade

// ProductInventoryDto 结构体
type ProductInventoryDto struct {
	// 日期 格式yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 库存数量
	InvCount int64 `json:"inv_count,omitempty" xml:"inv_count,omitempty"`
}
