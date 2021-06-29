package xhotelonlineorder

// UpdateOrderConfirmCodeParam 
type UpdateOrderConfirmCodeParam struct {

    // PMS确认号
    
    PmsResId   string `json:"pms_res_id,omitempty" xml:"pms_res_id,omitempty"`
    

    // 飞猪订单号
    
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    

    // 商家系统订单号
    
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    

}
