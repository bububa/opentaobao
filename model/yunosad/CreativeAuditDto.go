package yunosad

import (
	"sync"
)

// CreativeAuditDto 结构体
type CreativeAuditDto struct {
	// 广告创意id
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 是否审核通过，WAITING等待审核，PASS通过，REFUSE拒绝
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 拒绝原因
	RefuseCause string `json:"refuse_cause,omitempty" xml:"refuse_cause,omitempty"`
	// 创意级别
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolCreativeAuditDto = sync.Pool{
	New: func() any {
		return new(CreativeAuditDto)
	},
}

// GetCreativeAuditDto() 从对象池中获取CreativeAuditDto
func GetCreativeAuditDto() *CreativeAuditDto {
	return poolCreativeAuditDto.Get().(*CreativeAuditDto)
}

// ReleaseCreativeAuditDto 释放CreativeAuditDto
func ReleaseCreativeAuditDto(v *CreativeAuditDto) {
	v.CreativeId = ""
	v.Status = ""
	v.RefuseCause = ""
	v.Level = 0
	poolCreativeAuditDto.Put(v)
}
