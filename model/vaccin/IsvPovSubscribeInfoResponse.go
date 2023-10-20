package vaccin

import (
	"sync"
)

// IsvPovSubscribeInfoResponse 结构体
type IsvPovSubscribeInfoResponse struct {
	// 自有pov预约信息
	InfoDetailList []PovSubscribeDetailModel `json:"info_detail_list,omitempty" xml:"info_detail_list>pov_subscribe_detail_model,omitempty"`
	// 总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolIsvPovSubscribeInfoResponse = sync.Pool{
	New: func() any {
		return new(IsvPovSubscribeInfoResponse)
	},
}

// GetIsvPovSubscribeInfoResponse() 从对象池中获取IsvPovSubscribeInfoResponse
func GetIsvPovSubscribeInfoResponse() *IsvPovSubscribeInfoResponse {
	return poolIsvPovSubscribeInfoResponse.Get().(*IsvPovSubscribeInfoResponse)
}

// ReleaseIsvPovSubscribeInfoResponse 释放IsvPovSubscribeInfoResponse
func ReleaseIsvPovSubscribeInfoResponse(v *IsvPovSubscribeInfoResponse) {
	v.InfoDetailList = v.InfoDetailList[:0]
	v.Total = 0
	poolIsvPovSubscribeInfoResponse.Put(v)
}
