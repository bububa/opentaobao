package drugtrace

// AlibabaAlihealthDrugKytQuerydruginfoResultModel 
type AlibabaAlihealthDrugKytQuerydruginfoResultModel struct {
    // 返回业务对象
    Model   *CodeQueryDrugInfoDto `json:"model,omitempty" xml:"model,omitempty"`
    // 返回码
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
