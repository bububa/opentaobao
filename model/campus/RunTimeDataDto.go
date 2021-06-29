package campus

// RunTimeDataDTO 
type RunTimeDataDTO struct {
    // 设备id
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    // 事件时间
    EventTime   int64 `json:"event_time,omitempty" xml:"event_time,omitempty"`
    // 消息id
    MsgId   string `json:"msg_id,omitempty" xml:"msg_id,omitempty"`
    // runTimeData
    RunTimeDatas   []RunDataDTO `json:"run_time_datas,omitempty" xml:"run_time_datas>run_data_dto,omitempty"`
    // 是否是逻辑设备
    BeLogic   bool `json:"be_logic,omitempty" xml:"be_logic,omitempty"`
}
