package mozi

import (
	"sync"
)

// DismissOrganizationSupervisorResult 结构体
type DismissOrganizationSupervisorResult struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDismissOrganizationSupervisorResult = sync.Pool{
	New: func() any {
		return new(DismissOrganizationSupervisorResult)
	},
}

// GetDismissOrganizationSupervisorResult() 从对象池中获取DismissOrganizationSupervisorResult
func GetDismissOrganizationSupervisorResult() *DismissOrganizationSupervisorResult {
	return poolDismissOrganizationSupervisorResult.Get().(*DismissOrganizationSupervisorResult)
}

// ReleaseDismissOrganizationSupervisorResult 释放DismissOrganizationSupervisorResult
func ReleaseDismissOrganizationSupervisorResult(v *DismissOrganizationSupervisorResult) {
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseCode = ""
	v.Success = false
	poolDismissOrganizationSupervisorResult.Put(v)
}
