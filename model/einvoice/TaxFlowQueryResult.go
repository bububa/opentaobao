package einvoice

// TaxFlowQueryResult 
type TaxFlowQueryResult struct {
    // 入驻开通工单ID
    FlowId   string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
    // 入驻开通工单状态：  process: 入驻中；  success: 入驻成功；  closed: 入驻失败（发生退订或驳回工单）
    FlowStatus   string `json:"flow_status,omitempty" xml:"flow_status,omitempty"`
    // 开票商户信息
    InvoiceMerchant   *InvoiceMerchantDto `json:"invoice_merchant,omitempty" xml:"invoice_merchant,omitempty"`
    // 订购单信息
    InvoiceOrder   *InvoiceOrderSimpleDto `json:"invoice_order,omitempty" xml:"invoice_order,omitempty"`
}
