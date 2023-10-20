package miniappopen

import (
	"sync"
)

// DistributionOrderBindBaseDto 结构体
type DistributionOrderBindBaseDto struct {
	// 商品id
	TargetEntityId string `json:"target_entity_id,omitempty" xml:"target_entity_id,omitempty"`
	// 失败的原因
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 绑定是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDistributionOrderBindBaseDto = sync.Pool{
	New: func() any {
		return new(DistributionOrderBindBaseDto)
	},
}

// GetDistributionOrderBindBaseDto() 从对象池中获取DistributionOrderBindBaseDto
func GetDistributionOrderBindBaseDto() *DistributionOrderBindBaseDto {
	return poolDistributionOrderBindBaseDto.Get().(*DistributionOrderBindBaseDto)
}

// ReleaseDistributionOrderBindBaseDto 释放DistributionOrderBindBaseDto
func ReleaseDistributionOrderBindBaseDto(v *DistributionOrderBindBaseDto) {
	v.TargetEntityId = ""
	v.FailMsg = ""
	v.Success = false
	poolDistributionOrderBindBaseDto.Put(v)
}
