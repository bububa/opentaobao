package usergrowth2

import (
	"sync"
)

// CreativeAuditResultDto 结构体
type CreativeAuditResultDto struct {
	// 渠道-创意id
	OuterCreativeId string `json:"outer_creative_id,omitempty" xml:"outer_creative_id,omitempty"`
	// 审核结果
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 拒绝信息
	MaterialRejectInfo *MaterialRejectInfo `json:"material_reject_info,omitempty" xml:"material_reject_info,omitempty"`
}

var poolCreativeAuditResultDto = sync.Pool{
	New: func() any {
		return new(CreativeAuditResultDto)
	},
}

// GetCreativeAuditResultDto() 从对象池中获取CreativeAuditResultDto
func GetCreativeAuditResultDto() *CreativeAuditResultDto {
	return poolCreativeAuditResultDto.Get().(*CreativeAuditResultDto)
}

// ReleaseCreativeAuditResultDto 释放CreativeAuditResultDto
func ReleaseCreativeAuditResultDto(v *CreativeAuditResultDto) {
	v.OuterCreativeId = ""
	v.Status = ""
	v.MaterialRejectInfo = nil
	poolCreativeAuditResultDto.Put(v)
}
