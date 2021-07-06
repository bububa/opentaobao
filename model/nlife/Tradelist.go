package nlife

// Tradelist 结构体
type Tradelist struct {
	// 门店采购单列表
	TradeStoreList []Tradestorelist `json:"trade_store_list,omitempty" xml:"trade_store_list>tradestorelist,omitempty"`
	// 采购单生效时间
	GmtEffective string `json:"gmt_effective,omitempty" xml:"gmt_effective,omitempty"`
	// 企业采购单号
	EntTradeNo string `json:"ent_trade_no,omitempty" xml:"ent_trade_no,omitempty"`
}
