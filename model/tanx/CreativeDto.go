package tanx

import (
	"sync"
)

// CreativeDto 结构体
type CreativeDto struct {
	// 创意审核信息列表
	CreativePublisherAuditDtoList []CreativePublisherAuditDto `json:"creative_publisher_audit_dto_list,omitempty" xml:"creative_publisher_audit_dto_list>creative_publisher_audit_dto,omitempty"`
	// 创意ID
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意审核的状态（通过PASS,拒绝REFUSE,未审核WAITING）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 创意拒绝的原因
	RefuseCause string `json:"refuse_cause,omitempty" xml:"refuse_cause,omitempty"`
	// 广告位属性
	AdboardData string `json:"adboard_data,omitempty" xml:"adboard_data,omitempty"`
	// 创意通过的等级，1表示一级创意，99表示普通创意
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolCreativeDto = sync.Pool{
	New: func() any {
		return new(CreativeDto)
	},
}

// GetCreativeDto() 从对象池中获取CreativeDto
func GetCreativeDto() *CreativeDto {
	return poolCreativeDto.Get().(*CreativeDto)
}

// ReleaseCreativeDto 释放CreativeDto
func ReleaseCreativeDto(v *CreativeDto) {
	v.CreativePublisherAuditDtoList = v.CreativePublisherAuditDtoList[:0]
	v.CreativeId = ""
	v.Status = ""
	v.RefuseCause = ""
	v.AdboardData = ""
	v.Level = 0
	poolCreativeDto.Put(v)
}
