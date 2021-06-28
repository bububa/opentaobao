package tmallsc

// VerifyRequestDto 
/* model for simplify = false
type VerifyRequestDto struct {

    // 主订单id
    
    ParentOrderId   int64 `json:"parent_order_id,omitempty"`
    

    // 核销来源
    
    VerifySource   string `json:"verify_source,omitempty"`
    

    // 核销附加信息
    
    Attrs   string `json:"attrs,omitempty"`
    

    // 服务商id
    
    TpId   int64 `json:"tp_id,omitempty"`
    

    // 业务类型
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 核销码
    
    Code   string `json:"code,omitempty"`
    

    // 买家id
    
    BuyerId   int64 `json:"buyer_id,omitempty"`
    

    // 工单id
    
    WorkcardId   int64 `json:"workcard_id,omitempty"`
    

    // 子订单id
    
    OrderId   int64 `json:"order_id,omitempty"`
    

}
*/

// VerifyRequestDto 
type VerifyRequestDto struct {

    // 主订单id
    ParentOrderId   int64 `json:"parent_order_id,omitempty"`

    // 核销来源
    VerifySource   string `json:"verify_source,omitempty"`

    // 核销附加信息
    Attrs   string `json:"attrs,omitempty"`

    // 服务商id
    TpId   int64 `json:"tp_id,omitempty"`

    // 业务类型
    BizType   string `json:"biz_type,omitempty"`

    // 核销码
    Code   string `json:"code,omitempty"`

    // 买家id
    BuyerId   int64 `json:"buyer_id,omitempty"`

    // 工单id
    WorkcardId   int64 `json:"workcard_id,omitempty"`

    // 子订单id
    OrderId   int64 `json:"order_id,omitempty"`

}
