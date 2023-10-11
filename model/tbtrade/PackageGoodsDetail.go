package tbtrade

// PackageGoodsDetail 结构体
type PackageGoodsDetail struct {
	// 商品数字编号
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sku_id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
