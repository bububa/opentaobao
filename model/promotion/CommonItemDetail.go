package promotion

// CommonItemDetail 结构体
type CommonItemDetail struct {
	// 优惠活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 优惠详情ID
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 优惠类型，只有两种可选值：0-减钱；1-打折
	PromotionType int64 `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 优惠力度，其值的解释方式由promotion_type定义：当为减钱时解释成减钱数量，如：900表示减9元；当为打折时解释成打折折扣，如：900表示打9折
	PromotionValue int64 `json:"promotion_value,omitempty" xml:"promotion_value,omitempty"`
}
