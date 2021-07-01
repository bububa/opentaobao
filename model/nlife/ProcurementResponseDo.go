package nlife

// ProcurementResponseDo 结构体
type ProcurementResponseDo struct {
	// 企业ID
	EntId int64 `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 获取到的企业采购单的总数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 采购订单列表
	TradeList []Tradelist `json:"trade_list,omitempty" xml:"trade_list>tradelist,omitempty"`
}
