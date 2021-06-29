package feedflow

// TaobaoFeedflowItemTargetValidlistResultDTO 
type TaobaoFeedflowItemTargetValidlistResultDTO struct {
    // 成功
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 定向结构
    Targets   []TargetDTO `json:"targets,omitempty" xml:"targets>target_dto,omitempty"`
    // 调用是否成功,true-成功，false-失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
