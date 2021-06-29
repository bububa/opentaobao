package drugtrace

// AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel 
type AlibabaAlihealthDrugBillUpbillDetailWithcodeResultModel struct {

    // 最外层对象
    
    Model   *BillUpOutDetailDto `json:"model,omitempty" xml:"model,omitempty"`
    

    // 提示信息编码
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // 提示信息内容
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    

    // 成功失败标记
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
