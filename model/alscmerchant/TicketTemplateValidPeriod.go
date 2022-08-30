package alscmerchant

// TicketTemplateValidPeriod 结构体
type TicketTemplateValidPeriod struct {
	// 使用截止时间（绝对时间）(yyyy-MM-dd HH:mm:ss)
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 使用开始时间（绝对时间）(yyyy-MM-dd HH:mm:ss)
	ActiveTime string `json:"active_time,omitempty" xml:"active_time,omitempty"`
	// RELATIVE(相对类型) ABSOLUTE(绝对类型)
	ValidType string `json:"valid_type,omitempty" xml:"valid_type,omitempty"`
	// 使用截止相对分钟数（相对时间）
	EndMinTimes int64 `json:"end_min_times,omitempty" xml:"end_min_times,omitempty"`
	// 使用生效相对分钟数（相对时间）
	ActiveMinTimes int64 `json:"active_min_times,omitempty" xml:"active_min_times,omitempty"`
}
