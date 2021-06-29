package flight

// ServiceApiBean 
type ServiceApiBean struct {
    // 服务时间 1:不限制时间 2:按具体服务时间
    ServiceTimeType   int64 `json:"service_time_type,omitempty" xml:"service_time_type,omitempty"`
    // 服务内容。 最多允许200个字符。 禁止空格等特殊符号。
    ServiceContent   string `json:"service_content,omitempty" xml:"service_content,omitempty"`
    // 服务说明。 最多允许100个字符。 禁止空格等特殊符号。
    RoomTip   string `json:"room_tip,omitempty" xml:"room_tip,omitempty"`
    // 服务地点。 最多允许100个字符。 禁止空格等特殊符号。
    Room   string `json:"room,omitempty" xml:"room,omitempty"`
    // 服务时间结束 HHmm，serviceTimeType为2时必填
    ServiceTimeEnd   string `json:"service_time_end,omitempty" xml:"service_time_end,omitempty"`
    // 服务时间开始 HHmm，serviceTimeType为2时必填
    ServiceTimeStart   string `json:"service_time_start,omitempty" xml:"service_time_start,omitempty"`
}
