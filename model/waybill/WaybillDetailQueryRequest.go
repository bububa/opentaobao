package waybill

// WaybillDetailQueryRequest 结构体
type WaybillDetailQueryRequest struct {
	// 电子面单单号
	WaybillCodes []string `json:"waybill_codes,omitempty" xml:"waybill_codes>string,omitempty"`
	// CP快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 0:根据cp_code和waybil_code查询;1:根据订单号查询
	QueryBy int64 `json:"query_by,omitempty" xml:"query_by,omitempty"`
	// 交易订单号
	TradeOrderList []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
}
