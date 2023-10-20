package logistic

// Tms2mscPayQueryResponse 结构体
type Tms2mscPayQueryResponse struct {
	// 消费者支付状态
	PayFlag string `json:"pay_flag,omitempty" xml:"pay_flag,omitempty"`
	// 消费者支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
}
