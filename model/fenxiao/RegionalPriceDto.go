package fenxiao

// RegionalPriceDto 结构体
type RegionalPriceDto struct {
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 金额（分）
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 区县，特殊可选
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 街道，特殊可选
	Street string `json:"street,omitempty" xml:"street,omitempty"`
}
