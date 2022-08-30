package waybill

// ReachableRecommendResponseDto 结构体
type ReachableRecommendResponseDto struct {
	// 可达CP列表
	ReachableCps []CpInfoDto `json:"reachable_cps,omitempty" xml:"reachable_cps>cp_info_dto,omitempty"`
	// 不可达原因
	InterruptReason string `json:"interrupt_reason,omitempty" xml:"interrupt_reason,omitempty"`
	// 不可达结束时间
	InterruptEndTime string `json:"interrupt_end_time,omitempty" xml:"interrupt_end_time,omitempty"`
	// 拦截类型
	InterruptType string `json:"interrupt_type,omitempty" xml:"interrupt_type,omitempty"`
	// 地址是否可达
	Reachable bool `json:"reachable,omitempty" xml:"reachable,omitempty"`
}
