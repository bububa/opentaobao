package hotel

// RatePrice 结构体
type RatePrice struct {
	// 商品属性，INSTANT_CONFIRM(&#34;及时确认&#34;),MORNING_ORDER(&#34;支持凌晨入住&#34;),
	Attribute []string `json:"attribute,omitempty" xml:"attribute>string,omitempty"`
	// 日历报价列表
	CheckinCheckoutPrices []CheckInCheckOutPrice `json:"checkin_checkout_prices,omitempty" xml:"checkin_checkout_prices>check_in_check_out_price,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 每天可售起始时间，HH:mm:ss
	EffectiveDailyStartTime string `json:"effective_daily_start_time,omitempty" xml:"effective_daily_start_time,omitempty"`
	// 每天可售结束时间，HH:mm:ss
	EffectiveDailyEndTime string `json:"effective_daily_end_time,omitempty" xml:"effective_daily_end_time,omitempty"`
	// 可售开始时间，yyyy-MM-dd HH:mm:ss
	EffectiveStartTime string `json:"effective_start_time,omitempty" xml:"effective_start_time,omitempty"`
	// 可售结束时间，yyyy-MM-dd HH:mm:ss
	EffectiveEndTime string `json:"effective_end_time,omitempty" xml:"effective_end_time,omitempty"`
	// 退改政策json
	CancelPolicyJson string `json:"cancel_policy_json,omitempty" xml:"cancel_policy_json,omitempty"`
	// 退改政策描述信息
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// rateId
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// rpId
	RateplanId int64 `json:"rateplan_id,omitempty" xml:"rateplan_id,omitempty"`
	// iid
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 开票类型，0、不开票；1、门店开票；2、商家开票
	InvoiceProvider int64 `json:"invoice_provider,omitempty" xml:"invoice_provider,omitempty"`
	// 最大入住天数
	MaxStayDays int64 `json:"max_stay_days,omitempty" xml:"max_stay_days,omitempty"`
	// 最小入住天数
	MinStayDays int64 `json:"min_stay_days,omitempty" xml:"min_stay_days,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 最小提前预定小时数
	MinAdvanceHour int64 `json:"min_advance_hour,omitempty" xml:"min_advance_hour,omitempty"`
	// 最大提前预定小时数
	MaxAdvanceHour int64 `json:"max_advance_hour,omitempty" xml:"max_advance_hour,omitempty"`
}
