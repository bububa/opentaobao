package promotion

// BenefitReadTopQuery 结构体
type BenefitReadTopQuery struct {
	// 状态,online,offline,invalid
	Statuses []string `json:"statuses,omitempty" xml:"statuses>string,omitempty"`
	// 业务来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 权益发放开始时间大于该值
	StartTimeBegin string `json:"start_time_begin,omitempty" xml:"start_time_begin,omitempty"`
	// 权益发放开始时间小于该值
	StartTimeEnd string `json:"start_time_end,omitempty" xml:"start_time_end,omitempty"`
	// 权益发放结束时间小于该值
	EndTimeEnd string `json:"end_time_end,omitempty" xml:"end_time_end,omitempty"`
	// 权益发放结束时间大于该值
	EndTimeBegin string `json:"end_time_begin,omitempty" xml:"end_time_begin,omitempty"`
	// 权益类型
	BenefitType string `json:"benefit_type,omitempty" xml:"benefit_type,omitempty"`
	// 每页记录数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
