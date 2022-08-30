package product

// GoodsStatusDto 结构体
type GoodsStatusDto struct {
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}
