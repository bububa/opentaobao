package examination

// OrderInfo 
type OrderInfo struct {
    // 外部订单ID
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    // 订单状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 聊天地址，用户之间进行隔离
    ImUrl   string `json:"im_url,omitempty" xml:"im_url,omitempty"`
    // 医生ID
    DoctorId   string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
    // 订单ID
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
