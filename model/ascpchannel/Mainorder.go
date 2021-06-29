package ascpchannel

// Mainorder 
type Mainorder struct {

    // 操作id
    
    OperationOrderId   string `json:"operation_order_id,omitempty" xml:"operation_order_id,omitempty"`
    

    // 商家Uic_id
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

}
