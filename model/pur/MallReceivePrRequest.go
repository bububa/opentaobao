package pur

// MallReceivePrRequest 
type MallReceivePrRequest struct {

    // 第三方商城申请单ID
    
    PurReqId   string `json:"pur_req_id,omitempty" xml:"pur_req_id,omitempty"`
    

    // 用户ID
    
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

    // 商城编码
    
    AppCode   string `json:"app_code,omitempty" xml:"app_code,omitempty"`
    

    // 订单详情
    
    OrderItems   []OrderItem `json:"order_items,omitempty" xml:"order_items,omitempty"`
    

    // 附件信息
    
    MallFiles   []MallFile `json:"mall_files,omitempty" xml:"mall_files,omitempty"`
    

    // 外界商城申请标识
    
    OutPurReqId   string `json:"out_pur_req_id,omitempty" xml:"out_pur_req_id,omitempty"`
    

}
