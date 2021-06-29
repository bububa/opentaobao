package jipiao

// TaobaoAlitripSellerRefundGetResultDo 
type TaobaoAlitripSellerRefundGetResultDo struct {

    // 系统自动生成
    
    ErrorCode   string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
    

    // 系统自动生成
    
    ErrorMsg   string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
    

    // 申请单详情
    
    Results   *ReturnTicketDetail `json:"results,omitempty" xml:"results,omitempty"`
    

    // 调用是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

}
