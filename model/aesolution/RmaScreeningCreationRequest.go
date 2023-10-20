package aesolution

import (
	"sync"
)

// RmaScreeningCreationRequest 结构体
type RmaScreeningCreationRequest struct {
	// Values: OK, NO_OK
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// RMA ID
	RmaId string `json:"rma_id,omitempty" xml:"rma_id,omitempty"`
	// Date of screening
	ScreeningDate string `json:"screening_date,omitempty" xml:"screening_date,omitempty"`
	// Description of the screening result
	ScreeningResultDetails string `json:"screening_result_details,omitempty" xml:"screening_result_details,omitempty"`
	// Values: CUSTOMER_FAULT, GIVE_UP_UNSEALED, GIVE_UP_SEALED, DOA_SEALED_QUALITY_ISSUE, DOA_SEALED_NO_QUALITY_ISSUE
	ScreeningResultReasons string `json:"screening_result_reasons,omitempty" xml:"screening_result_reasons,omitempty"`
}

var poolRmaScreeningCreationRequest = sync.Pool{
	New: func() any {
		return new(RmaScreeningCreationRequest)
	},
}

// GetRmaScreeningCreationRequest() 从对象池中获取RmaScreeningCreationRequest
func GetRmaScreeningCreationRequest() *RmaScreeningCreationRequest {
	return poolRmaScreeningCreationRequest.Get().(*RmaScreeningCreationRequest)
}

// ReleaseRmaScreeningCreationRequest 释放RmaScreeningCreationRequest
func ReleaseRmaScreeningCreationRequest(v *RmaScreeningCreationRequest) {
	v.Result = ""
	v.RmaId = ""
	v.ScreeningDate = ""
	v.ScreeningResultDetails = ""
	v.ScreeningResultReasons = ""
	poolRmaScreeningCreationRequest.Put(v)
}
