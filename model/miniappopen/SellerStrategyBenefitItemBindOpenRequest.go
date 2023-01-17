package miniappopen

// SellerStrategyBenefitItemBindOpenRequest 结构体
type SellerStrategyBenefitItemBindOpenRequest struct {
	// C小程序id
	AppId int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 策略中，实物权益对应的奖池id
	PoolId int64 `json:"pool_id,omitempty" xml:"pool_id,omitempty"`
}
