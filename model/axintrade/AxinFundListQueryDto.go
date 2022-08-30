package axintrade

// AxinFundListQueryDto 结构体
type AxinFundListQueryDto struct {
	// 资金单类型
	FundStatus []string `json:"fund_status,omitempty" xml:"fund_status>string,omitempty"`
	// 外部订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
}
