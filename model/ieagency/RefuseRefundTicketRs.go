package ieagency

import (
	"sync"
)

// RefuseRefundTicketRs 结构体
type RefuseRefundTicketRs struct {
	// apiErrorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// apiErrorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRefuseRefundTicketRs = sync.Pool{
	New: func() any {
		return new(RefuseRefundTicketRs)
	},
}

// GetRefuseRefundTicketRs() 从对象池中获取RefuseRefundTicketRs
func GetRefuseRefundTicketRs() *RefuseRefundTicketRs {
	return poolRefuseRefundTicketRs.Get().(*RefuseRefundTicketRs)
}

// ReleaseRefuseRefundTicketRs 释放RefuseRefundTicketRs
func ReleaseRefuseRefundTicketRs(v *RefuseRefundTicketRs) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Success = false
	poolRefuseRefundTicketRs.Put(v)
}
