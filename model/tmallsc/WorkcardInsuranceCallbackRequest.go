package tmallsc

import (
	"sync"
)

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

var poolWorkcardInsuranceCallbackRequest = sync.Pool{
	New: func() any {
		return new(WorkcardInsuranceCallbackRequest)
	},
}

// GetWorkcardInsuranceCallbackRequest() 从对象池中获取WorkcardInsuranceCallbackRequest
func GetWorkcardInsuranceCallbackRequest() *WorkcardInsuranceCallbackRequest {
	return poolWorkcardInsuranceCallbackRequest.Get().(*WorkcardInsuranceCallbackRequest)
}

// ReleaseWorkcardInsuranceCallbackRequest 释放WorkcardInsuranceCallbackRequest
func ReleaseWorkcardInsuranceCallbackRequest(v *WorkcardInsuranceCallbackRequest) {
	v.ClaimDesc = ""
	v.ClaimOrderNo = ""
	v.ClaimTime = ""
	v.WorkcardId = 0
	v.ClaimCount = 0
	v.ClaimStatus = 0
	v.ClaimFee = 0
	poolWorkcardInsuranceCallbackRequest.Put(v)
}
