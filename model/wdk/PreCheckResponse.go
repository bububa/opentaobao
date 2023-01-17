package wdk

// PreCheckResponse 结构体
type PreCheckResponse struct {
	// 商品是否可作业信息
	ProductList []PreCheckProductResponse `json:"product_list,omitempty" xml:"product_list>pre_check_product_response,omitempty"`
	// 履约前置校验扩展数据
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
}
