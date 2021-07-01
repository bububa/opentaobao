package scbp

// TopContextDto 结构体
type TopContextDto struct {
	// 产品线id
	ProductLineId int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
