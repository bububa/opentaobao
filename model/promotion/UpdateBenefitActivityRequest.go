package promotion

import (
	"sync"
)

// UpdateBenefitActivityRequest 结构体
type UpdateBenefitActivityRequest struct {
	// 活动关联的权益信息，可以从权益选择器API中获取
	AddDetailVos []ActivityBenefitDetailVo `json:"add_detail_vos,omitempty" xml:"add_detail_vos>activity_benefit_detail_vo,omitempty"`
	// 需要删除的已经关联的权益
	DeleteDetailVos []DeleteActivityBenefitDetailVo `json:"delete_detail_vos,omitempty" xml:"delete_detail_vos>delete_activity_benefit_detail_vo,omitempty"`
	// 同步权益活动的概述信息，方便卖家后台查看
	BenefitActivityVo *UpdateBenefitActivityVo `json:"benefit_activity_vo,omitempty" xml:"benefit_activity_vo,omitempty"`
}

var poolUpdateBenefitActivityRequest = sync.Pool{
	New: func() any {
		return new(UpdateBenefitActivityRequest)
	},
}

// GetUpdateBenefitActivityRequest() 从对象池中获取UpdateBenefitActivityRequest
func GetUpdateBenefitActivityRequest() *UpdateBenefitActivityRequest {
	return poolUpdateBenefitActivityRequest.Get().(*UpdateBenefitActivityRequest)
}

// ReleaseUpdateBenefitActivityRequest 释放UpdateBenefitActivityRequest
func ReleaseUpdateBenefitActivityRequest(v *UpdateBenefitActivityRequest) {
	v.AddDetailVos = v.AddDetailVos[:0]
	v.DeleteDetailVos = v.DeleteDetailVos[:0]
	v.BenefitActivityVo = nil
	poolUpdateBenefitActivityRequest.Put(v)
}
