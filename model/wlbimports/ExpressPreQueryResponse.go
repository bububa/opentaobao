package wlbimports

// ExpressPreQueryResponse 结构体
type ExpressPreQueryResponse struct {
	// 产品code
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 揽收地址当地截单时间
	LocalCutoffDateAndTime string `json:"local_cutoff_date_and_time,omitempty" xml:"local_cutoff_date_and_time,omitempty"`
	// 预计送达时间
	EstimatedDeliveryTime string `json:"estimated_delivery_time,omitempty" xml:"estimated_delivery_time,omitempty"`
	// 最早揽收时间
	PickupEarliest string `json:"pickup_earliest,omitempty" xml:"pickup_earliest,omitempty"`
	// 最晚揽收时间
	PickupLatest string `json:"pickup_latest,omitempty" xml:"pickup_latest,omitempty"`
	// 运输总天数
	TotalTransitDays int64 `json:"total_transit_days,omitempty" xml:"total_transit_days,omitempty"`
}
