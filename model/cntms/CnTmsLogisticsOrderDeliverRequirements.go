package cntms

// CnTmsLogisticsOrderDeliverRequirements 
type CnTmsLogisticsOrderDeliverRequirements struct {

    // 配送类型： PTPS-普通配送 LLPS-冷链配送
    DeliveryType   string `json:"delivery_type,omitempty"`

    // 投递时延要求(1 工作日 2 节假日 104 预约达 )
    ScheduleType   string `json:"schedule_type,omitempty"`

    // 送达日期（格式为 yyyy-MM-dd)
    ScheduleDay   string `json:"schedule_day,omitempty"`

    // 送达开始时间（格式为 hh:mm）
    ScheduleStart   string `json:"schedule_start,omitempty"`

    // 送达结束时间（格式为 hh:mm）
    ScheduleEnd   string `json:"schedule_end,omitempty"`

}
