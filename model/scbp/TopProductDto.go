package scbp

// TopProductDto 结构体
type TopProductDto struct {
	// 产品标题，最大长度256个字符
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 产品推广状态，取值[disabled,enabled]
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
