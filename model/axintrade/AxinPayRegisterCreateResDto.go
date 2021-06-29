package axintrade

// AxinPayRegisterCreateResDto 
type AxinPayRegisterCreateResDto struct {

    // 支付宝返回的申请单号
    
    ApplyOrderId   string `json:"apply_order_id,omitempty" xml:"apply_order_id,omitempty"`
    

    // 支付平台申请单号
    
    PayRegisterOrderNo   string `json:"pay_register_order_no,omitempty" xml:"pay_register_order_no,omitempty"`
    

}
