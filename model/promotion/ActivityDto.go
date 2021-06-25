package promotion

// ActivityDto 
type ActivityDto struct {

    // 权益列表
    Benefits   []BenefitDto `json:"benefits,omitempty"`

    // 活动来源
    Source   string `json:"source,omitempty"`

    // 活动来源记录id
    SourceRecordId   string `json:"source_record_id,omitempty"`

    // 活动扩展数据
    Feature   string `json:"feature,omitempty"`

    // 活动名称
    Name   string `json:"name,omitempty"`

    // 活动投放code
    StrategyCode   string `json:"strategy_code,omitempty"`

    // 活动开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 活动结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 活动状态
    Status   string `json:"status,omitempty"`

    // 活动投放渠道
    ChannelCode   string `json:"channel_code,omitempty"`

}
