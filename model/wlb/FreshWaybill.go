package wlb

// FreshWaybill 
/* model for simplify = false
type FreshWaybill struct {

    // 获取的所有电子面单号，以“;”分隔
    
    WaybillCode   string `json:"waybill_code,omitempty"`
    

    // 大头笔
    
    ShortAddress   string `json:"short_address,omitempty"`
    

    // 简称
    
    Alias   string `json:"alias,omitempty"`
    

    // 预计到达时间
    
    Time   string `json:"time,omitempty"`
    

    // 预留扩展字段
    
    Feature   string `json:"feature,omitempty"`
    

    // 交易号
    
    TradeId   string `json:"trade_id,omitempty"`
    

}
*/

// FreshWaybill 
type FreshWaybill struct {

    // 获取的所有电子面单号，以“;”分隔
    WaybillCode   string `json:"waybill_code,omitempty"`

    // 大头笔
    ShortAddress   string `json:"short_address,omitempty"`

    // 简称
    Alias   string `json:"alias,omitempty"`

    // 预计到达时间
    Time   string `json:"time,omitempty"`

    // 预留扩展字段
    Feature   string `json:"feature,omitempty"`

    // 交易号
    TradeId   string `json:"trade_id,omitempty"`

}
