package tmallsc

// WorkcardInsuranceCallbackRequest 结构体
type WorkcardInsuranceCallbackRequest struct {
	// 拒绝理赔原因（拒绝时启用
	ClaimDesc string `json:"claim_desc,omitempty" xml:"claim_desc,omitempty"`
	// 理赔单号
	ClaimOrderNo string `json:"claim_order_no,omitempty" xml:"claim_order_no,omitempty"`
	// 理赔时间
	ClaimTime string `json:"claim_time,omitempty" xml:"claim_time,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 理赔单数量：理论都为1
	ClaimCount int64 `json:"claim_count,omitempty" xml:"claim_count,omitempty"`
	// 理赔状态0：未理赔 1 理赔成功 2：理赔失败
	ClaimStatus int64 `json:"claim_status,omitempty" xml:"claim_status,omitempty"`
	// 理赔金额（分
	ClaimFee int64 `json:"claim_fee,omitempty" xml:"claim_fee,omitempty"`
}
