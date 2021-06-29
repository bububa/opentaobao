package jipiao

// TaobaoAlitripSellerRefundSearchResultDO 
type TaobaoAlitripSellerRefundSearchResultDO struct {
    // 错误码
    ErrorCode   string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
    // ReturnTicketDo
    Results   []ReturnTicketDO `json:"results,omitempty" xml:"results>return_ticket_do,omitempty"`
    // 调用是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
