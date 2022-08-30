package logistic

// CommodityInfo 结构体
type CommodityInfo struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品数量
	GoodsQuantity int64 `json:"goods_quantity,omitempty" xml:"goods_quantity,omitempty"`
}
