package tanx

import (
	"sync"
)

// CreativePublisherAuditDto 结构体
type CreativePublisherAuditDto struct {
	// 状态（PASS.通过；REFUSE.拒绝；WAITING.待审）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 媒体ID
	PubliserId string `json:"publiser_id,omitempty" xml:"publiser_id,omitempty"`
	// 拒绝原因
	RefuseCause string `json:"refuse_cause,omitempty" xml:"refuse_cause,omitempty"`
}

var poolCreativePublisherAuditDto = sync.Pool{
	New: func() any {
		return new(CreativePublisherAuditDto)
	},
}

// GetCreativePublisherAuditDto() 从对象池中获取CreativePublisherAuditDto
func GetCreativePublisherAuditDto() *CreativePublisherAuditDto {
	return poolCreativePublisherAuditDto.Get().(*CreativePublisherAuditDto)
}

// ReleaseCreativePublisherAuditDto 释放CreativePublisherAuditDto
func ReleaseCreativePublisherAuditDto(v *CreativePublisherAuditDto) {
	v.Status = ""
	v.PubliserId = ""
	v.RefuseCause = ""
	poolCreativePublisherAuditDto.Put(v)
}
