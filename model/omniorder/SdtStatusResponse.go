package omniorder

// SdtStatusResponse 
type SdtStatusResponse struct {

    // 卖家ID，通sellerID
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // packageId
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 取消原因
    
    ReasonDesc   string `json:"reason_desc,omitempty" xml:"reason_desc,omitempty"`
    

    // 状态 0 取号，1 已发货 -1 商家取消 -2 运力端取消
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

}
