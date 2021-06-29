package ascpchannel

// ExternalRefundGoodsWaybillRequest 
type ExternalRefundGoodsWaybillRequest struct {

    // 外部退款单号
    
    OutRefundNo   string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
    

    // 退款单号
    
    RefundNo   string `json:"refund_no,omitempty" xml:"refund_no,omitempty"`
    

    // 物流公司名称
    
    LogisticsCompanyName   string `json:"logistics_company_name,omitempty" xml:"logistics_company_name,omitempty"`
    

    // 物流公司编码
    
    LogisticsCompanyCode   string `json:"logistics_company_code,omitempty" xml:"logistics_company_code,omitempty"`
    

    // 物流单号
    
    LogisticsWaybillNo   string `json:"logistics_waybill_no,omitempty" xml:"logistics_waybill_no,omitempty"`
    

}
