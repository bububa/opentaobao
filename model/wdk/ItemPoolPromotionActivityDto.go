package wdk

// ItemPoolPromotionActivityDto 结构体
type ItemPoolPromotionActivityDto struct {
	// 活动周几生效
	Weekdays []string `json:"weekdays,omitempty" xml:"weekdays>string,omitempty"`
	// 活动每天生效时间段
	EveryDayPeriods []string `json:"every_day_periods,omitempty" xml:"every_day_periods>string,omitempty"`
	// 优惠适用场景:1:APP  2:POS
	Terminals []string `json:"terminals,omitempty" xml:"terminals>string,omitempty"`
	// 门店列表
	StoreIds []string `json:"store_ids,omitempty" xml:"store_ids>string,omitempty"`
	// 外部门店列表
	OuterStoreIds []string `json:"outer_store_ids,omitempty" xml:"outer_store_ids>string,omitempty"`
	// 人群编码,saas平台人群编码:NEW_USER: 新用户, OLD_USER：老用户，LIGHT_NEW_USER：闪购新客
	MemberCrowdCodes []string `json:"member_crowd_codes,omitempty" xml:"member_crowd_codes>string,omitempty"`
	// 逻辑分组
	LogicGroups []LogicGroupDto `json:"logic_groups,omitempty" xml:"logic_groups>logic_group_dto,omitempty"`
	// 阶梯分组
	StairGroups []StairGroupDto `json:"stair_groups,omitempty" xml:"stair_groups>stair_group_dto,omitempty"`
	// 商品池活动类型 AMOUNT_DECREASE("amount_decrease", "满元减"),     AMOUNT_DISCOUNT("amount_discount", "满元折"),     COUNT_DECREASE("count_decrease", "满件减"),     COUNT_DISCOUNT("count_discount", "满件折"),     COUNT_N_FIXED_PRICE("count_n_fixed_price", "满N件第N件一口价"),     COUNT_N_DISCOUNT("count_n_discount", "满N件第N件Y折"),     AMOUNT_EXCHANGE("amount_exchange", "满元换购"),     COUNT_EXCHANGE("count_exchange", "满件换购"),     COUNT_ALL_FIXED_PRICE("count_all_fixed_price", "X元Y件"),
	ItemPoolDiscountType string `json:"item_pool_discount_type,omitempty" xml:"item_pool_discount_type,omitempty"`
	// 外部订单号
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 操作人ID
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人姓名
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 营销活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
	// 活动开始时间
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 活动结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 限购信息
	Limit *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
	// 是否上不封顶
	EnableMultiple bool `json:"enable_multiple,omitempty" xml:"enable_multiple,omitempty"`
}
