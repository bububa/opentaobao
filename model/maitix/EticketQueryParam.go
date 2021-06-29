package maitix

// EticketQueryParam 
type EticketQueryParam struct {

    // 主分销单Id 必传
    
    MainOrderId   int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
    

    // 子分销单Id 可传可不传
    
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    

}
