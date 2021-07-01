package servicecenter

// OfnPreRedPacketFundRecordDto 结构体
type OfnPreRedPacketFundRecordDto struct {
	// 资产编号
	FundId int64 `json:"fund_id,omitempty" xml:"fund_id,omitempty"`
	// 变化金额
	ChangeAmount int64 `json:"change_amount,omitempty" xml:"change_amount,omitempty"`
}
