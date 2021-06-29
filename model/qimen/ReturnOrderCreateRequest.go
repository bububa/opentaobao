package qimen

// ReturnOrderCreateRequest 
type ReturnOrderCreateRequest struct {

    // 退货单信息
    
    ReturnOrder   *ReturnOrder `json:"returnOrder,omitempty" xml:"returnOrder,omitempty"`
    

    // 订单信息
    
    OrderLines   []OrderLine `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *TaobaoQimenReturnorderCreateMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

}
