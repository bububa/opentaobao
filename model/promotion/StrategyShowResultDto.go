package promotion

// StrategyShowResultDto 
type StrategyShowResultDto struct {

    // 扩展参数
    ExtraData   string `json:"extra_data,omitempty"`

    // 是否有下一页
    HasNextPage   bool `json:"has_next_page,omitempty"`

    // 投放计划
    ShowStrategy   *ShowStrategyDto `json:"show_strategy,omitempty"`

    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 权益列表
    ShowBenefits   []ShowBenefitDto `json:"show_benefits,omitempty"`

    // 追踪信息
    TrackingData   string `json:"tracking_data,omitempty"`

}
