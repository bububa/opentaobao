package wdk

// DeliveryCallbackOrderDetail 
/* model for simplify = false
type DeliveryCallbackOrderDetail struct {

    // 子单ID
    
    WorkOrderDetailId   string `json:"work_order_detail_id,omitempty"`
    

    // 拒收原因
    
    RefusedReason   string `json:"refused_reason,omitempty"`
    

}
*/

// DeliveryCallbackOrderDetail 
type DeliveryCallbackOrderDetail struct {

    // 子单ID
    WorkOrderDetailId   string `json:"work_order_detail_id,omitempty"`

    // 拒收原因
    RefusedReason   string `json:"refused_reason,omitempty"`

}
