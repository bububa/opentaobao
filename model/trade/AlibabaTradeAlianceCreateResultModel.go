package trade

// AlibabaTradeAlianceCreateResultModel 
type AlibabaTradeAlianceCreateResultModel struct {

    // 是否创建成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 生成的订单号i
    
    MainOrderId   string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    

    // totalAmount
    
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    

}
