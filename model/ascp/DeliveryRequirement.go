package ascp

// DeliveryRequirement 结构体
type DeliveryRequirement struct {
	// 投递时延要求(1=工作日;2=节假日;101=当日达;102=次晨达;103=次日达;104=预约达;105=隔日达)
	ScheduleType string `json:"schedule_type,omitempty" xml:"schedule_type,omitempty"`
	// 要求送达日期(YYYY-MM-DD)
	ScheduleDay string `json:"schedule_day,omitempty" xml:"schedule_day,omitempty"`
	// 投递时间范围要求(开始时间;格式：HH:MM:SS)
	ScheduleStartTime string `json:"schedule_start_time,omitempty" xml:"schedule_start_time,omitempty"`
	// 投递时间范围要求(结束时间;格式：HH:MM:SS)
	ScheduleEndTime string `json:"schedule_end_time,omitempty" xml:"schedule_end_time,omitempty"`
	// 发货服务类型(PTPS:普通配送;LLPS:冷链配送;HBP:环保配)
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}
