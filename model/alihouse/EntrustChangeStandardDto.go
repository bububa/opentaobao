package alihouse

import (
	"sync"
)

// EntrustChangeStandardDto 结构体
type EntrustChangeStandardDto struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 房源E码
	Ecode string `json:"ecode,omitempty" xml:"ecode,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 委托房源外部房源id
	EntrustOuterId string `json:"entrust_outer_id,omitempty" xml:"entrust_outer_id,omitempty"`
	// 存在跨小区委托变标准的场景，那么该字段是标准房源的小区id
	PublicCommunityOuterId string `json:"public_community_outer_id,omitempty" xml:"public_community_outer_id,omitempty"`
	// 业务类型，1-新房，2-二手房，3-租房，默认为2
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 0-正常房源 1-临时房源（交易专用）
	HouseType int64 `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 租房业务模式
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
}

var poolEntrustChangeStandardDto = sync.Pool{
	New: func() any {
		return new(EntrustChangeStandardDto)
	},
}

// GetEntrustChangeStandardDto() 从对象池中获取EntrustChangeStandardDto
func GetEntrustChangeStandardDto() *EntrustChangeStandardDto {
	return poolEntrustChangeStandardDto.Get().(*EntrustChangeStandardDto)
}

// ReleaseEntrustChangeStandardDto 释放EntrustChangeStandardDto
func ReleaseEntrustChangeStandardDto(v *EntrustChangeStandardDto) {
	v.CommunityOuterId = ""
	v.Ecode = ""
	v.OuterId = ""
	v.EntrustOuterId = ""
	v.PublicCommunityOuterId = ""
	v.BusinessType = 0
	v.HouseType = 0
	v.HouseModel = 0
	poolEntrustChangeStandardDto.Put(v)
}
