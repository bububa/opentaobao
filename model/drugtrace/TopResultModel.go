package drugtrace

// TopResultModel 
type TopResultModel struct {
    // 顶层结构
    Model   *AdvanceCodeSearchDTO `json:"model,omitempty" xml:"model,omitempty"`
    // 操作码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 操作说明
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 调用成功
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
    // 导出的项目和药品目录
    Models   []TrialProjectDTO `json:"models,omitempty" xml:"models>trial_project_dto,omitempty"`
}
