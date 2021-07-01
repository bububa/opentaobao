package qianniu

// OrderStatisticsResult 结构体
type OrderStatisticsResult struct {
	// tqdj_order_num
	TqdjOrderNum int64 `json:"tqdj_order_num,omitempty" xml:"tqdj_order_num,omitempty"`
	// take_order_num
	TakeOrderNum int64 `json:"take_order_num,omitempty" xml:"take_order_num,omitempty"`
	// home_order_num
	HomeOrderNum int64 `json:"home_order_num,omitempty" xml:"home_order_num,omitempty"`
	// step_order_num
	StepOrderNum int64 `json:"step_order_num,omitempty" xml:"step_order_num,omitempty"`
}
