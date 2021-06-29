package einvoice

// InvoiceFlowRefundDto 
type InvoiceFlowRefundDto struct {

    // 退款备注信息
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 工单id(入驻、续约、加盘)
    
    FlowId   string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
    

    // 业务平台码
    
    PlatformCode   string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
    

}
