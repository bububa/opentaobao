package aliospay

import (
	"sync"
)

// GetPeriodAgreementStatusResponse 结构体
type GetPeriodAgreementStatusResponse struct {
	// 周期扣款协议签约id
	AgreementId string `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 周期扣款协议签约状态，INIT-初始，SIGNED-已签约，UNSIGNED-已解约
	AgreementStatus string `json:"agreement_status,omitempty" xml:"agreement_status,omitempty"`
	// 协议签约时间，时间戳，单位毫秒
	SignTime int64 `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	// 协议解约时间，时间戳，单位毫秒
	UnsignTime int64 `json:"unsign_time,omitempty" xml:"unsign_time,omitempty"`
}

var poolGetPeriodAgreementStatusResponse = sync.Pool{
	New: func() any {
		return new(GetPeriodAgreementStatusResponse)
	},
}

// GetGetPeriodAgreementStatusResponse() 从对象池中获取GetPeriodAgreementStatusResponse
func GetGetPeriodAgreementStatusResponse() *GetPeriodAgreementStatusResponse {
	return poolGetPeriodAgreementStatusResponse.Get().(*GetPeriodAgreementStatusResponse)
}

// ReleaseGetPeriodAgreementStatusResponse 释放GetPeriodAgreementStatusResponse
func ReleaseGetPeriodAgreementStatusResponse(v *GetPeriodAgreementStatusResponse) {
	v.AgreementId = ""
	v.AgreementStatus = ""
	v.SignTime = 0
	v.UnsignTime = 0
	poolGetPeriodAgreementStatusResponse.Put(v)
}
