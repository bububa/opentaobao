package moscm

// CspuInputDto 结构体
type CspuInputDto struct {
	// 销售属性
	Properties []PropertyDto `json:"properties,omitempty" xml:"properties>property_dto,omitempty"`
	// 货号
	ArtNo string `json:"art_no,omitempty" xml:"art_no,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 源id（供应商自己的唯一标识）
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 吊牌价
	TagPrice string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
	// brand是银泰品牌Id，colorName（颜色名称）、colorCode(颜色编码)、sizeCode(尺码编码)、sizeName(尺码名称):商品销售属性，标签一些在属性里找不到id的属性存放在这里
	Tags string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 产品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 标准商品信息（款）
	SpuInputDto *SpuInputDto `json:"spu_input_dto,omitempty" xml:"spu_input_dto,omitempty"`
	// 天猫Sku
	TmallSkuId int64 `json:"tmall_sku_id,omitempty" xml:"tmall_sku_id,omitempty"`
}
