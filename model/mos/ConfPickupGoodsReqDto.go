package mos

// ConfPickupGoodsReqDto 
/* model for simplify = false
type ConfPickupGoodsReqDto struct {

    // OC交易号
    
    TradeNo   string `json:"trade_no,omitempty"`
    

    // 核销码
    
    ShortCode   string `json:"short_code,omitempty"`
    

    // 操作员
    
    Operator   string `json:"operator,omitempty"`
    

    // 门店号
    
    OutBelongId   string `json:"out_belong_id,omitempty"`
    

    // 专柜号
    
    OutStoreId   string `json:"out_store_id,omitempty"`
    

}
*/

// ConfPickupGoodsReqDto 
type ConfPickupGoodsReqDto struct {

    // OC交易号
    TradeNo   string `json:"trade_no,omitempty"`

    // 核销码
    ShortCode   string `json:"short_code,omitempty"`

    // 操作员
    Operator   string `json:"operator,omitempty"`

    // 门店号
    OutBelongId   string `json:"out_belong_id,omitempty"`

    // 专柜号
    OutStoreId   string `json:"out_store_id,omitempty"`

}
