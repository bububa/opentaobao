package axindata

// FscRouteProductAddResponse 结构体
type FscRouteProductAddResponse struct {
	// 产品ID
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 原始产品ID
	OriProductId string `json:"ori_product_id,omitempty" xml:"ori_product_id,omitempty"`
	// 产品编码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
}
