package drugtrace

import (
	"sync"
)

// PUserEntInfoDto 结构体
type PUserEntInfoDto struct {
	// 所在地编码
	DictRegionCode string `json:"dict_region_code,omitempty" xml:"dict_region_code,omitempty"`
	// 所在地明细
	RegRegionDetail string `json:"reg_region_detail,omitempty" xml:"reg_region_detail,omitempty"`
	// 注册地编码
	RegRegionCode string `json:"reg_region_code,omitempty" xml:"reg_region_code,omitempty"`
	// 注册地明细
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 是否法人
	LegalOrgFlag string `json:"legal_org_flag,omitempty" xml:"legal_org_flag,omitempty"`
	// 所属管理机构
	DirectManage string `json:"direct_manage,omitempty" xml:"direct_manage,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 是否入网
	IsNetwork string `json:"is_network,omitempty" xml:"is_network,omitempty"`
	// 企业唯一标识
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 企业id
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 所在地明细
	DictRegionDetail string `json:"dict_region_detail,omitempty" xml:"dict_region_detail,omitempty"`
	// 状态1.使用中0.已废除
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 拼音缩写
	EntCapitalName string `json:"ent_capital_name,omitempty" xml:"ent_capital_name,omitempty"`
	// 企业类型
	UserRoleType string `json:"user_role_type,omitempty" xml:"user_role_type,omitempty"`
	// 企业类型编码
	UserRoleTypeStr string `json:"user_role_type_str,omitempty" xml:"user_role_type_str,omitempty"`
	// 企业机构详细类别
	EntOrgType string `json:"ent_org_type,omitempty" xml:"ent_org_type,omitempty"`
	// 省
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 市
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 县
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 入网标识【0非入网1入网2入驻马上放心】
	NetworkType string `json:"network_type,omitempty" xml:"network_type,omitempty"`
}

var poolPUserEntInfoDto = sync.Pool{
	New: func() any {
		return new(PUserEntInfoDto)
	},
}

// GetPUserEntInfoDto() 从对象池中获取PUserEntInfoDto
func GetPUserEntInfoDto() *PUserEntInfoDto {
	return poolPUserEntInfoDto.Get().(*PUserEntInfoDto)
}

// ReleasePUserEntInfoDto 释放PUserEntInfoDto
func ReleasePUserEntInfoDto(v *PUserEntInfoDto) {
	v.DictRegionCode = ""
	v.RegRegionDetail = ""
	v.RegRegionCode = ""
	v.OrgCode = ""
	v.LegalOrgFlag = ""
	v.DirectManage = ""
	v.EntName = ""
	v.IsNetwork = ""
	v.RefEntId = ""
	v.EntId = ""
	v.DictRegionDetail = ""
	v.Status = ""
	v.EntCapitalName = ""
	v.UserRoleType = ""
	v.UserRoleTypeStr = ""
	v.EntOrgType = ""
	v.ProvName = ""
	v.AreaName = ""
	v.CityName = ""
	v.NetworkType = ""
	poolPUserEntInfoDto.Put(v)
}
