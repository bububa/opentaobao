package tblogistics

// AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult 结构体
type AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否已取消
	Canceled bool `json:"canceled,omitempty" xml:"canceled,omitempty"`
}
