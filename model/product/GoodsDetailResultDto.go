package product

// GoodsDetailResultDto 结构体
type GoodsDetailResultDto struct {
	// 商品详情信息
	GoodsDetail *ExternalGoodsDetailDto `json:"goods_detail,omitempty" xml:"goods_detail,omitempty"`
	// 商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}
