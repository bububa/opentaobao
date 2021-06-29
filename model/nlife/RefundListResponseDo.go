package nlife

// RefundListResponseDo 
type RefundListResponseDo struct {

    // 获取到的结果的总数量
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
    

    // 采购退货单列表
    
    TradeRefundList   []Traderefundlist `json:"trade_refund_list,omitempty" xml:"trade_refund_list,omitempty"`
    

    // 企业Id
    
    EntId   int64 `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
    

}
