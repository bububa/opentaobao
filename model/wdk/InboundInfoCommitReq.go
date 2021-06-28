package wdk

// InboundInfoCommitReq 
type InboundInfoCommitReq struct {

    // 入库商品明细
    
    InboundItemInfos   []InboundItemInfo `json:"inbound_item_infos,omitempty" xml:"inbound_item_infos,omitempty"`
    

    // 收货入库单号
    
    InboundOrderNo   string `json:"inbound_order_no,omitempty" xml:"inbound_order_no,omitempty"`
    

    // 收货时间
    
    ReceivedTime   string `json:"received_time,omitempty" xml:"received_time,omitempty"`
    

    // 采购退货单单号
    
    ReturnOrderNo   string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    

}
