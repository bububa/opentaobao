package tmallgenie

import (
	"sync"
)

// VipCycleThirdVipStatusRequest 结构体
type VipCycleThirdVipStatusRequest struct {
	// 三方用户id
	ThirdUserId string `json:"third_user_id,omitempty" xml:"third_user_id,omitempty"`
	// 三方源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 合约类型：1-包月 2-包季 3-包年
	PerformType int64 `json:"perform_type,omitempty" xml:"perform_type,omitempty"`
}

var poolVipCycleThirdVipStatusRequest = sync.Pool{
	New: func() any {
		return new(VipCycleThirdVipStatusRequest)
	},
}

// GetVipCycleThirdVipStatusRequest() 从对象池中获取VipCycleThirdVipStatusRequest
func GetVipCycleThirdVipStatusRequest() *VipCycleThirdVipStatusRequest {
	return poolVipCycleThirdVipStatusRequest.Get().(*VipCycleThirdVipStatusRequest)
}

// ReleaseVipCycleThirdVipStatusRequest 释放VipCycleThirdVipStatusRequest
func ReleaseVipCycleThirdVipStatusRequest(v *VipCycleThirdVipStatusRequest) {
	v.ThirdUserId = ""
	v.Source = ""
	v.PerformType = 0
	poolVipCycleThirdVipStatusRequest.Put(v)
}
