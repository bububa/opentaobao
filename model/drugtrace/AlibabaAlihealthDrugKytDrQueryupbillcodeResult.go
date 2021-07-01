package drugtrace

// AlibabaAlihealthDrugKytDrQueryupbillcodeResult 
type AlibabaAlihealthDrugKytDrQueryupbillcodeResult struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // model
    ModelList   []BillUpstreamDto `json:"model_list,omitempty" xml:"model_list>bill_upstream_dto,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
