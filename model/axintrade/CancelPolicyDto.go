package axintrade

import (
	"sync"
)

// CancelPolicyDto 结构体
type CancelPolicyDto struct {
	// 详细规则
	PolicyInfoList []CancelPolicyInfoDto `json:"policy_info_list,omitempty" xml:"policy_info_list>cancel_policy_info_dto,omitempty"`
	// 取消政策类型
	CancelPolicyType int64 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

var poolCancelPolicyDto = sync.Pool{
	New: func() any {
		return new(CancelPolicyDto)
	},
}

// GetCancelPolicyDto() 从对象池中获取CancelPolicyDto
func GetCancelPolicyDto() *CancelPolicyDto {
	return poolCancelPolicyDto.Get().(*CancelPolicyDto)
}

// ReleaseCancelPolicyDto 释放CancelPolicyDto
func ReleaseCancelPolicyDto(v *CancelPolicyDto) {
	v.PolicyInfoList = v.PolicyInfoList[:0]
	v.CancelPolicyType = 0
	poolCancelPolicyDto.Put(v)
}
