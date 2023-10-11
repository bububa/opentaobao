package tblogistics

// TopConsignGoodsRequest 结构体
type TopConsignGoodsRequest struct {
	// 子订单id
	SubTid string `json:"sub_tid,omitempty" xml:"sub_tid,omitempty"`
	// 成分品itemId
	CompItemId string `json:"comp_item_id,omitempty" xml:"comp_item_id,omitempty"`
	// 成分品skuId
	CompSkuId string `json:"comp_sku_id,omitempty" xml:"comp_sku_id,omitempty"`
	// 商品类型 0：普通品 1:赠品 2:成分品，默认0
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 商品数量，不传默认为子单上的商品数量；支持不传，但不能传0或负值
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
