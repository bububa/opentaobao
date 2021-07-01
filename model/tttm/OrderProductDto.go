package tttm

// OrderProductDto 结构体
type OrderProductDto struct {
	// 货品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 货品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 货品数量
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 货品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
}
