package alihealth2

// AlibabaAlihealthStoreCertificateCreateResult 
type AlibabaAlihealthStoreCertificateCreateResult struct {

    // 状态
    
    Status   bool `json:"status,omitempty" xml:"status,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 创建审批流是否成功
    
    Module   bool `json:"module,omitempty" xml:"module,omitempty"`
    

}
