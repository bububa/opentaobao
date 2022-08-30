package tmallsc

// ServiceProduct 结构体
type ServiceProduct struct {
	// 服务产品名称
	ProductTitle string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// 服务产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
