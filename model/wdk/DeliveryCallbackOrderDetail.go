package wdk

// DeliveryCallbackOrderDetail 
type DeliveryCallbackOrderDetail struct {

    // 子单ID
    
    WorkOrderDetailId   string `json:"work_order_detail_id,omitempty" xml:"work_order_detail_id,omitempty"`
    

    // 拒收原因
    
    RefusedReason   string `json:"refused_reason,omitempty" xml:"refused_reason,omitempty"`
    

}
