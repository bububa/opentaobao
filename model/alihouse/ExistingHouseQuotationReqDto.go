package alihouse

import (
	"sync"
)

// ExistingHouseQuotationReqDto 结构体
type ExistingHouseQuotationReqDto struct {
	// 发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 目标id
	TargetId string `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// 行情内容
	ContentJson string `json:"content_json,omitempty" xml:"content_json,omitempty"`
	// 业务类型
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 是否测试
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 毫秒级时间戳
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
}

var poolExistingHouseQuotationReqDto = sync.Pool{
	New: func() any {
		return new(ExistingHouseQuotationReqDto)
	},
}

// GetExistingHouseQuotationReqDto() 从对象池中获取ExistingHouseQuotationReqDto
func GetExistingHouseQuotationReqDto() *ExistingHouseQuotationReqDto {
	return poolExistingHouseQuotationReqDto.Get().(*ExistingHouseQuotationReqDto)
}

// ReleaseExistingHouseQuotationReqDto 释放ExistingHouseQuotationReqDto
func ReleaseExistingHouseQuotationReqDto(v *ExistingHouseQuotationReqDto) {
	v.PublishTime = ""
	v.TargetId = ""
	v.ContentJson = ""
	v.BizType = 0
	v.IsTest = 0
	v.Type = 0
	v.Status = 0
	v.EtcVersion = 0
	poolExistingHouseQuotationReqDto.Put(v)
}
