package drugtrace

import (
	"sync"
)

// PUserEntDto 结构体
type PUserEntDto struct {
	// 机构编码
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 原企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 新企业名称
	EntNameNew string `json:"ent_name_new,omitempty" xml:"ent_name_new,omitempty"`
	// 企业refid
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 企业id
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
}

var poolPUserEntDto = sync.Pool{
	New: func() any {
		return new(PUserEntDto)
	},
}

// GetPUserEntDto() 从对象池中获取PUserEntDto
func GetPUserEntDto() *PUserEntDto {
	return poolPUserEntDto.Get().(*PUserEntDto)
}

// ReleasePUserEntDto 释放PUserEntDto
func ReleasePUserEntDto(v *PUserEntDto) {
	v.OrgCode = ""
	v.EntName = ""
	v.EntNameNew = ""
	v.RefEntId = ""
	v.EntId = ""
	poolPUserEntDto.Put(v)
}
