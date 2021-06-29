package axintrade

// AxinFundListQueryDTO 
type AxinFundListQueryDTO struct {
    // 资金单类型
    FundStatus   []int64 `json:"fund_status,omitempty" xml:"fund_status>int64,omitempty"`
    // 外部订单号
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
}
