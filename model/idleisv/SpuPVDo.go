package idleisv

// SpuPVDo 结构体
type SpuPVDo struct {
	// 品牌和型号信息
	BrandModelList []IdleNewPubValueDo `json:"brand_model_list,omitempty" xml:"brand_model_list>idle_new_pub_value_do,omitempty"`
	// 品牌型号的显示值
	SpuTitle string `json:"spu_title,omitempty" xml:"spu_title,omitempty"`
}
