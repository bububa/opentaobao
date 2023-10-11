package jym

// JymSingleGoodsDto 结构体
type JymSingleGoodsDto struct {
	// 标签
	Tags []JymGoodsTagDto `json:"tags,omitempty" xml:"tags>jym_goods_tag_dto,omitempty"`
	// 商品ID
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 商品标题
	GoodsTitle string `json:"goods_title,omitempty" xml:"goods_title,omitempty"`
	// 商品头图地址
	HeaderImgUrl string `json:"header_img_url,omitempty" xml:"header_img_url,omitempty"`
	// 商品单价
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}
