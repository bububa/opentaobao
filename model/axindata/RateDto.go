package axindata

// RateDto 结构体
type RateDto struct {
	// 价库日历
	PriceStockDtoList []PriceStockDto `json:"price_stock_dto_list,omitempty" xml:"price_stock_dto_list>price_stock_dto,omitempty"`
	// 早餐信息列表
	BreakFastList []BreakfastDto `json:"break_fast_list,omitempty" xml:"break_fast_list>breakfast_dto,omitempty"`
	// 取消规则列表
	CancelRuleList []CancelPolicyRuleDto `json:"cancel_rule_list,omitempty" xml:"cancel_rule_list>cancel_policy_rule_dto,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 售卖政策名称
	RatePlanName string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty"`
	// 资源标签
	ResourceTag string `json:"resource_tag,omitempty" xml:"resource_tag,omitempty"`
	// 佣金费率
	CommissionFeeRate string `json:"commission_fee_rate,omitempty" xml:"commission_fee_rate,omitempty"`
	// 每日截止时间
	EndTimeDaily string `json:"end_time_daily,omitempty" xml:"end_time_daily,omitempty"`
	// 是否即时确认
	InstantConfirm string `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
	// 价格计划英文名
	RatePlanNameEn string `json:"rate_plan_name_en,omitempty" xml:"rate_plan_name_en,omitempty"`
	// 每日开始时间
	StartTimeDaily string `json:"start_time_daily,omitempty" xml:"start_time_daily,omitempty"`
	// 分销模式 FP-底价模式 SP-卖家模式
	DistributeMode string `json:"distribute_mode,omitempty" xml:"distribute_mode,omitempty"`
	// 当前rate修改时间，精确到毫秒
	ModifiedTime int64 `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// rateId
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 售卖政策id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 最小提前小时数
	MinAdvHours int64 `json:"min_adv_hours,omitempty" xml:"min_adv_hours,omitempty"`
	// 最大提前小时数
	MaxAdvHours int64 `json:"max_adv_hours,omitempty" xml:"max_adv_hours,omitempty"`
	// 最小入住天数
	MinStay int64 `json:"min_stay,omitempty" xml:"min_stay,omitempty"`
	// 最大入住天数
	MaxStay int64 `json:"max_stay,omitempty" xml:"max_stay,omitempty"`
	// 最小预订间数
	MinBookCount int64 `json:"min_book_count,omitempty" xml:"min_book_count,omitempty"`
	// 最大预订间数
	MaxBookCount int64 `json:"max_book_count,omitempty" xml:"max_book_count,omitempty"`
	// 连住天数,适用于复杂价格
	Nod int64 `json:"nod,omitempty" xml:"nod,omitempty"`
	// 入住人数,适用于复杂价格
	Nop int64 `json:"nop,omitempty" xml:"nop,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 最大连住天数
	MaxDays int64 `json:"max_days,omitempty" xml:"max_days,omitempty"`
	// 凌晨房信息
	DawnBookingDto *DawnBookingDto `json:"dawn_booking_dto,omitempty" xml:"dawn_booking_dto,omitempty"`
	// 是否复杂价格
	MultiplePrice bool `json:"multiple_price,omitempty" xml:"multiple_price,omitempty"`
}
