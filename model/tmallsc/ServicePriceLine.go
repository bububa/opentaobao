package tmallsc

// ServicePriceLine 结构体
type ServicePriceLine struct {
	// 服务外部编码
	ServiceOuterId string `json:"service_outer_id,omitempty" xml:"service_outer_id,omitempty"`
	// 原价
	ConsumerPrice int64 `json:"consumer_price,omitempty" xml:"consumer_price,omitempty"`
	// 教育优惠价格
	EduPrice int64 `json:"edu_price,omitempty" xml:"edu_price,omitempty"`
}
