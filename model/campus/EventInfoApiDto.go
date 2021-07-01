package campus

// EventInfoApiDto 
type EventInfoApiDto struct {
    // 事件时间
    EventTime   string `json:"event_time,omitempty" xml:"event_time,omitempty"`
    // 事件信息ID
    EventInfoId   int64 `json:"event_info_id,omitempty" xml:"event_info_id,omitempty"`
    // 规则实例code
    RuleInstanceCode   string `json:"rule_instance_code,omitempty" xml:"rule_instance_code,omitempty"`
    // IBos报警事件状态上报
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 事件类型code
    EventTypeCode   string `json:"event_type_code,omitempty" xml:"event_type_code,omitempty"`
}
