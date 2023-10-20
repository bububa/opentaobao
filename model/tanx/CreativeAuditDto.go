package tanx

import (
	"sync"
)

// CreativeAuditDto 结构体
type CreativeAuditDto struct {
	// 媒体审核列表
	PublishersAuditInfoList []CreativePublisherAuditDto `json:"publishers_audit_info_list,omitempty" xml:"publishers_audit_info_list>creative_publisher_audit_dto,omitempty"`
	// 创意ID
	CreativeId string `json:"creative_id,omitempty" xml:"creative_id,omitempty"`
	// 创意审核的状态（通过PASS,拒绝REFUSE,未审核WAITING）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 创意拒绝的原因
	RefuseCause string `json:"refuse_cause,omitempty" xml:"refuse_cause,omitempty"`
	// 创意内容数据
	AdboardData string `json:"adboard_data,omitempty" xml:"adboard_data,omitempty"`
	// 广告主ID
	AdvertiserIds string `json:"advertiser_ids,omitempty" xml:"advertiser_ids,omitempty"`
	// 创意通过的等级，1表示一级创意，99表示普通创意
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// DSP用户ID
	DspId int64 `json:"dsp_id,omitempty" xml:"dsp_id,omitempty"`
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
	v.PublishersAuditInfoList = v.PublishersAuditInfoList[:0]
	v.CreativeId = ""
	v.Status = ""
	v.RefuseCause = ""
	v.AdboardData = ""
	v.AdvertiserIds = ""
	v.Level = 0
	v.DspId = 0
	poolCreativeAuditDto.Put(v)
}
