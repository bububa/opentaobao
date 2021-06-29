package alihealthcrm

// BabyInfoResult 
type BabyInfoResult struct {
    // 操作码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 操作说明
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 宝宝基本信息
    Model   *BabyBaseInfoDTO `json:"model,omitempty" xml:"model,omitempty"`
}
