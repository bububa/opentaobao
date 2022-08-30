package alihouse

// SyncTradePosApplyDto 结构体
type SyncTradePosApplyDto struct {
	// 经营主体id
	MerchantOpenId string `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// pos数量
	PosApplyAmount string `json:"pos_apply_amount,omitempty" xml:"pos_apply_amount,omitempty"`
	// pos进件id
	PosOpenId string `json:"pos_open_id,omitempty" xml:"pos_open_id,omitempty"`
	// 通联id
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
}
