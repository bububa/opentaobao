package drugtrace

// AlibabaAlihealthDrugKytCodetobillResultModel 
type AlibabaAlihealthDrugKytCodetobillResultModel struct {
    // 返回结果
    Model   *CodeToBill `json:"model,omitempty" xml:"model,omitempty"`
    // 调用编码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 调用结果
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回结果
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
