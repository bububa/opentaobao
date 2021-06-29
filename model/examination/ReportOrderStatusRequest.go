package examination

// ReportOrderStatusRequest 
type ReportOrderStatusRequest struct {
    // 备注
    Note   string `json:"note,omitempty" xml:"note,omitempty"`
    // 订单状态，10=待接诊，20=已关闭，30=问诊中，60=已结束
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 订单ID
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 外部订单ID
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    // 医生ID
    DoctorId   string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
}
