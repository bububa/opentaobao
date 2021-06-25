package qimen

// DeliveryRequirements 
type DeliveryRequirements struct {

    // 投递时延要求(1=工作日;2=节假日;101=当日达;102=次晨达;103=次日达;104= 预约 达)
    ScheduleType   int64 `json:"scheduleType,omitempty"`

    // 要求送达日期(YYYY-MM-DD)
    ScheduleDay   string `json:"scheduleDay,omitempty"`

    // 投递时间范围要求(开始时间;格式：HH:MM:SS)
    ScheduleStartTime   string `json:"scheduleStartTime,omitempty"`

    // 投递时间范围要求(结束时间;格式：HH:MM:SS)
    ScheduleEndTime   string `json:"scheduleEndTime,omitempty"`

    // 发货服务类型(PTPS:普通配送;LLPS:冷链配送;HBP:环保配)
    DeliveryType   string `json:"deliveryType,omitempty"`

}
