package nlife

// Tradelist 
type Tradelist struct {

    // 采购单生效时间
    
    GmtEffective   string `json:"gmt_effective,omitempty" xml:"gmt_effective,omitempty"`
    

    // 企业采购单号
    
    EntTradeNo   string `json:"ent_trade_no,omitempty" xml:"ent_trade_no,omitempty"`
    

    // 门店采购单列表
    
    TradeStoreList   []Tradestorelist `json:"trade_store_list,omitempty" xml:"trade_store_list,omitempty"`
    

}
