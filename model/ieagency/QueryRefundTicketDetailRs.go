package ieagency

import (
	"sync"
)

// QueryRefundTicketDetailRs 结构体
type QueryRefundTicketDetailRs struct {
	// apiErrorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// ieRefundTicketVO
	RefundTicket *IeRefundTicketVo `json:"refund_ticket,omitempty" xml:"refund_ticket,omitempty"`
	// apiErrorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolQueryRefundTicketDetailRs = sync.Pool{
	New: func() any {
		return new(QueryRefundTicketDetailRs)
	},
}

// GetQueryRefundTicketDetailRs() 从对象池中获取QueryRefundTicketDetailRs
func GetQueryRefundTicketDetailRs() *QueryRefundTicketDetailRs {
	return poolQueryRefundTicketDetailRs.Get().(*QueryRefundTicketDetailRs)
}

// ReleaseQueryRefundTicketDetailRs 释放QueryRefundTicketDetailRs
func ReleaseQueryRefundTicketDetailRs(v *QueryRefundTicketDetailRs) {
	v.ErrorMsg = ""
	v.RefundTicket = nil
	v.ErrorCode = 0
	v.Success = false
	poolQueryRefundTicketDetailRs.Put(v)
}
