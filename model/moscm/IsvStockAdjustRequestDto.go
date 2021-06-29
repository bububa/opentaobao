package moscm

// IsvStockAdjustRequestDto 
type IsvStockAdjustRequestDto struct {

    // 入库项（最大列表长度：20）
    
    InboundItems   []IsvInboundRequestItemDto `json:"inbound_items,omitempty" xml:"inbound_items,omitempty"`
    

    // 外部单号
    
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    

    // 出库项（最大列表长度：20）
    
    OutboundItems   []IsvOutboundRequestItemDto `json:"outbound_items,omitempty" xml:"outbound_items,omitempty"`
    

    // 备注
    
    Remarks   string `json:"remarks,omitempty" xml:"remarks,omitempty"`
    

    // 专柜Id
    
    CounterId   string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
    

    // 外部专柜Id（两个专柜id字段至少填写一个，如果同时填写，已counter_id为准）
    
    OutCounterId   string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
    

}
