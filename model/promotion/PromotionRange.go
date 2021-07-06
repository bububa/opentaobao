package promotion

// PromotionRange 结构体
type PromotionRange struct {
	// 活动名称。
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 参与活动的商品id。
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 活动id。
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
