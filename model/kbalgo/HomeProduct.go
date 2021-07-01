package kbalgo

// HomeProduct 结构体
type HomeProduct struct {
	// 到家基本信息
	BaseInfo *BaseInfo `json:"base_info,omitempty" xml:"base_info,omitempty"`
	// 商品信息
	ProductInfos []ProductInfo `json:"product_infos,omitempty" xml:"product_infos>product_info,omitempty"`
}
