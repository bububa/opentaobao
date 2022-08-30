package product

// GoodsPriceDto 结构体
type GoodsPriceDto struct {
	// 商品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 外部商品id对象
	ExternalGoodsId *ExternalGoodsIdDto `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
}
