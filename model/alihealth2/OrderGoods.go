package alihealth2

// OrderGoods 结构体
type OrderGoods struct {
	// 商品实际购买价
	RealPrice string `json:"real_price,omitempty" xml:"real_price,omitempty"`
	// 商品ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 商品购买数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 商品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商品外部ID
	GoodsCode string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
	// 商品ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
