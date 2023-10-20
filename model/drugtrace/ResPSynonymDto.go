package drugtrace

import (
	"sync"
)

// ResPSynonymDto 结构体
type ResPSynonymDto struct {
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 市
	CityDesc string `json:"city_desc,omitempty" xml:"city_desc,omitempty"`
	// 省
	ProvDesc string `json:"prov_desc,omitempty" xml:"prov_desc,omitempty"`
	// 区
	AreaDesc string `json:"area_desc,omitempty" xml:"area_desc,omitempty"`
	// 企业主键
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 区域编码
	DictRegionCode string `json:"dict_region_code,omitempty" xml:"dict_region_code,omitempty"`
	// 企业唯一标识
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 货主
	SynOwnEntId string `json:"syn_own_ent_id,omitempty" xml:"syn_own_ent_id,omitempty"`
	// 货主标识
	UserEntId string `json:"user_ent_id,omitempty" xml:"user_ent_id,omitempty"`
	// 创建日期
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 角色
	UserRoleType string `json:"user_role_type,omitempty" xml:"user_role_type,omitempty"`
}

var poolResPSynonymDto = sync.Pool{
	New: func() any {
		return new(ResPSynonymDto)
	},
}

// GetResPSynonymDto() 从对象池中获取ResPSynonymDto
func GetResPSynonymDto() *ResPSynonymDto {
	return poolResPSynonymDto.Get().(*ResPSynonymDto)
}

// ReleaseResPSynonymDto 释放ResPSynonymDto
func ReleaseResPSynonymDto(v *ResPSynonymDto) {
	v.EntName = ""
	v.CityDesc = ""
	v.ProvDesc = ""
	v.AreaDesc = ""
	v.EntId = ""
	v.DictRegionCode = ""
	v.RefEntId = ""
	v.SynOwnEntId = ""
	v.UserEntId = ""
	v.CrtDate = ""
	v.UserRoleType = ""
	poolResPSynonymDto.Put(v)
}
