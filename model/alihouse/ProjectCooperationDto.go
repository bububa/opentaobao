package alihouse

// ProjectCooperationDto 
type ProjectCooperationDto struct {
    // 合作开始时间
    CooperationEndTime   string `json:"cooperation_end_time,omitempty" xml:"cooperation_end_time,omitempty"`
    // 合作结束时间
    CooperationBeginTime   string `json:"cooperation_begin_time,omitempty" xml:"cooperation_begin_time,omitempty"`
    // KA楼盘ID
    KaProjectMid   string `json:"ka_project_mid,omitempty" xml:"ka_project_mid,omitempty"`
    // 流水号
    OuterCooperationId   string `json:"outer_cooperation_id,omitempty" xml:"outer_cooperation_id,omitempty"`
    // 外部楼盘ID
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 状态 0-无效 1-有效
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 是否优先展示 0-否 1-是
    IsPriority   int64 `json:"is_priority,omitempty" xml:"is_priority,omitempty"`
}
