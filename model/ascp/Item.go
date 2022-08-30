package ascp

// Item 结构体
type Item struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 商品单价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品仓储系统编码
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 单据行号
	OrderLineNo string `json:"order_line_no,omitempty" xml:"order_line_no,omitempty"`
}
