package tmallservice

// SettlementPriceFactor 结构体
type SettlementPriceFactor struct {
	// 计价因子属性
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 计价因子说明
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 计价因子实际值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}
