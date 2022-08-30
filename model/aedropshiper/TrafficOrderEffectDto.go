package aedropshiper

// TrafficOrderEffectDto 结构体
type TrafficOrderEffectDto struct {
	// commission rate
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// finished amount of the order, unit cent
	FinishedAmount string `json:"finished_amount,omitempty" xml:"finished_amount,omitempty"`
	// is affiliate product
	IsAffiliateProduct string `json:"is_affiliate_product,omitempty" xml:"is_affiliate_product,omitempty"`
	// is new buyer
	IsNewBuyer string `json:"is_new_buyer,omitempty" xml:"is_new_buyer,omitempty"`
	// item title
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// Additional order status, eg: full refund order, antispam order
	EffectDetailStatus string `json:"effect_detail_status,omitempty" xml:"effect_detail_status,omitempty"`
	// estimated commission for finished incentive order
	EstimatedIncentiveFinishedCommission string `json:"estimated_incentive_finished_commission,omitempty" xml:"estimated_incentive_finished_commission,omitempty"`
	// estimated commission for paid incentive order
	EstimatedIncentivePaidCommission string `json:"estimated_incentive_paid_commission,omitempty" xml:"estimated_incentive_paid_commission,omitempty"`
	// is hot product
	IsHotProduct string `json:"is_hot_product,omitempty" xml:"is_hot_product,omitempty"`
	// item detail url
	ItemDetailUrl string `json:"item_detail_url,omitempty" xml:"item_detail_url,omitempty"`
	// created time of this order
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// publisher settled currency
	PublisherSettledCurrency string `json:"publisher_settled_currency,omitempty" xml:"publisher_settled_currency,omitempty"`
	// product shipping country
	ShipToCountry string `json:"ship_to_country,omitempty" xml:"ship_to_country,omitempty"`
	// item main image url
	ItemMainImageUrl string `json:"item_main_image_url,omitempty" xml:"item_main_image_url,omitempty"`
	// paid time
	PaidTime string `json:"paid_time,omitempty" xml:"paid_time,omitempty"`
	// item count
	ItemCount string `json:"item_count,omitempty" xml:"item_count,omitempty"`
	// effect status
	EffectStatus string `json:"effect_status,omitempty" xml:"effect_status,omitempty"`
	// estimated commission for finished order
	EstimatedFinishedCommission string `json:"estimated_finished_commission,omitempty" xml:"estimated_finished_commission,omitempty"`
	// Order finish time
	FinishedTime string `json:"finished_time,omitempty" xml:"finished_time,omitempty"`
	// incentive commission rate
	IncentiveCommissionRate string `json:"incentive_commission_rate,omitempty" xml:"incentive_commission_rate,omitempty"`
	// order number
	OrderNumber int64 `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// publisher id
	PublisherId int64 `json:"publisher_id,omitempty" xml:"publisher_id,omitempty"`
	// parent order number
	ParentOrderNumber int64 `json:"parent_order_number,omitempty" xml:"parent_order_number,omitempty"`
	// order id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// item id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sub order id
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// estimated commission for paid order
	EstimatedPaidCommission int64 `json:"estimated_paid_commission,omitempty" xml:"estimated_paid_commission,omitempty"`
	// payment amount of the order, unit cent
	PaidAmount int64 `json:"paid_amount,omitempty" xml:"paid_amount,omitempty"`
	// category id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
