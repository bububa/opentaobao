package drugtrace

// AlibabaAlihealthDrugKytStorebilllistResultModel 
type AlibabaAlihealthDrugKytStorebilllistResultModel struct {
    // 返回结果
    Model   *PageInfoDTO `json:"model,omitempty" xml:"model,omitempty"`
    // 返回码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 返回是否成功
    ResponseSuccess   bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
