package alsc

// QueryRechargeAccountFlowOpenReq 
type QueryRechargeAccountFlowOpenReq struct {

    // 品牌ID(不能和outbrandid同时为空)
    
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 流水Id
    
    FlowId   string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
    

    // 外部品牌ID(不能和brandid同时为空)
    
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    

    // 外部订单ID，和流水ID必传一
    
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    

}
