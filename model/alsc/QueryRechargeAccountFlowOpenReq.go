package alsc

// QueryRechargeAccountFlowOpenReq 
/* model for simplify = false
type QueryRechargeAccountFlowOpenReq struct {

    // 品牌ID(不能和outbrandid同时为空)
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 流水Id
    
    FlowId   string `json:"flow_id,omitempty"`
    

    // 外部品牌ID(不能和brandid同时为空)
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // 外部订单ID，和流水ID必传一
    
    OuterOrderId   string `json:"outer_order_id,omitempty"`
    

}
*/

// QueryRechargeAccountFlowOpenReq 
type QueryRechargeAccountFlowOpenReq struct {

    // 品牌ID(不能和outbrandid同时为空)
    BrandId   string `json:"brand_id,omitempty"`

    // 流水Id
    FlowId   string `json:"flow_id,omitempty"`

    // 外部品牌ID(不能和brandid同时为空)
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // 外部订单ID，和流水ID必传一
    OuterOrderId   string `json:"outer_order_id,omitempty"`

}
