package feedflow

// TaobaoFeedflowItemCrowdPageResultDTO 
type TaobaoFeedflowItemCrowdPageResultDTO struct {
    // 成功
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 人群列表
    Crowds   []CrowdDTO `json:"crowds,omitempty" xml:"crowds>crowd_dto,omitempty"`
    // 总数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 调用是否成功,true-成功，false-失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
