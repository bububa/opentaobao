package ascp

import (
	"sync"
)

// CancelDistributeRequest 结构体
type CancelDistributeRequest struct {
	// 取消详情
	CancelDistributeInfoList []CancelDistributeInfo `json:"cancel_distribute_info_list,omitempty" xml:"cancel_distribute_info_list>cancel_distribute_info,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolCancelDistributeRequest = sync.Pool{
	New: func() any {
		return new(CancelDistributeRequest)
	},
}

// GetCancelDistributeRequest() 从对象池中获取CancelDistributeRequest
func GetCancelDistributeRequest() *CancelDistributeRequest {
	return poolCancelDistributeRequest.Get().(*CancelDistributeRequest)
}

// ReleaseCancelDistributeRequest 释放CancelDistributeRequest
func ReleaseCancelDistributeRequest(v *CancelDistributeRequest) {
	v.CancelDistributeInfoList = v.CancelDistributeInfoList[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolCancelDistributeRequest.Put(v)
}
