package xhotelonlineorder

// TaobaoXhotelCommoninvoiceListVtwoResultSet 
type TaobaoXhotelCommoninvoiceListVtwoResultSet struct {

    // 是否成功标记
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 常用发票信息
    
    Results   []CommonInvoiceInfo `json:"results,omitempty" xml:"results,omitempty"`
    

}
