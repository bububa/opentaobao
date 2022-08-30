package ascp

// Package 结构体
type Package struct {
	// 实际发货的配编码
	LogisticsCode string `json:"logistics_code,omitempty" xml:"logistics_code,omitempty"`
	// 运单号
	ExpressCode string `json:"express_code,omitempty" xml:"express_code,omitempty"`
	// 货品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 发货数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
