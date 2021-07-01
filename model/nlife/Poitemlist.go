package nlife

// Poitemlist 结构体
type Poitemlist struct {
	// 商品itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 采购数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 采购价-单位:分
	PoPrice int64 `json:"po_price,omitempty" xml:"po_price,omitempty"`
	// 零售价-单位:分
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 扣点-百分数换算成double的值
	ShareRatio string `json:"share_ratio,omitempty" xml:"share_ratio,omitempty"`
	// 活动扣点-百分数换算成double的值
	ActivityShareRatio string `json:"activity_share_ratio,omitempty" xml:"activity_share_ratio,omitempty"`
}
