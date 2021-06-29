package drugtrace

// AlibabaAlihealthDrugKytSearchbillResultModel 
type AlibabaAlihealthDrugKytSearchbillResultModel struct {
    // 返回模型
    Model   *PageInfoDTO `json:"model,omitempty" xml:"model,omitempty"`
    // 返回码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 是否响应成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
