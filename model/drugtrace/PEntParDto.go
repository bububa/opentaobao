package drugtrace

import (
	"sync"
)

// PEntParDto 结构体
type PEntParDto struct {
	// 往来单位ID
	PartnerId string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 往来单位名称
	PartnerName string `json:"partner_name,omitempty" xml:"partner_name,omitempty"`
	// 企业id
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 往来单位企业唯一标识
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 往来单位企业所在省编码
	EntProvCode string `json:"ent_prov_code,omitempty" xml:"ent_prov_code,omitempty"`
	// 所在省
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 所在市
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 所在县
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 是不是入网企业
	IsNetwork string `json:"is_network,omitempty" xml:"is_network,omitempty"`
	// 拼音缩写
	PartnerCapitalName string `json:"partner_capital_name,omitempty" xml:"partner_capital_name,omitempty"`
	// 类型
	PartnerType string `json:"partner_type,omitempty" xml:"partner_type,omitempty"`
	// 往来单位企业id
	PartnerEntId string `json:"partner_ent_id,omitempty" xml:"partner_ent_id,omitempty"`
	// 最近修改日期
	LastModDate string `json:"last_mod_date,omitempty" xml:"last_mod_date,omitempty"`
	// 创建日期
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 创建IC名称
	CrtIcName string `json:"crt_ic_name,omitempty" xml:"crt_ic_name,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 修改IC名称
	ModIcName string `json:"mod_ic_name,omitempty" xml:"mod_ic_name,omitempty"`
	// 级别
	PartnerLevel string `json:"partner_level,omitempty" xml:"partner_level,omitempty"`
	// 修改IC码
	ModIcCode string `json:"mod_ic_code,omitempty" xml:"mod_ic_code,omitempty"`
	// 合作ID
	PEntParId string `json:"p_ent_par_id,omitempty" xml:"p_ent_par_id,omitempty"`
	// 创建IC码
	CrtIcCode string `json:"crt_ic_code,omitempty" xml:"crt_ic_code,omitempty"`
	// 往来单位企业ID
	ParRefEntId string `json:"par_ref_ent_id,omitempty" xml:"par_ref_ent_id,omitempty"`
	// 证件类型名称
	LicenseTypeStr string `json:"license_type_str,omitempty" xml:"license_type_str,omitempty"`
	// 详细地址
	AddrDetail string `json:"addr_detail,omitempty" xml:"addr_detail,omitempty"`
	// 城市
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 1-审核通过；2-审核不通过；0-待审核：服务端针对networkType和standard字段汇总给出状态
	AuditFlag int64 `json:"audit_flag,omitempty" xml:"audit_flag,omitempty"`
	// networkType
	NetworkType int64 `json:"network_type,omitempty" xml:"network_type,omitempty"`
}

var poolPEntParDto = sync.Pool{
	New: func() any {
		return new(PEntParDto)
	},
}

// GetPEntParDto() 从对象池中获取PEntParDto
func GetPEntParDto() *PEntParDto {
	return poolPEntParDto.Get().(*PEntParDto)
}

// ReleasePEntParDto 释放PEntParDto
func ReleasePEntParDto(v *PEntParDto) {
	v.PartnerId = ""
	v.PartnerName = ""
	v.EntId = ""
	v.RefEntId = ""
	v.EntProvCode = ""
	v.ProvName = ""
	v.AreaName = ""
	v.CityName = ""
	v.IsNetwork = ""
	v.PartnerCapitalName = ""
	v.PartnerType = ""
	v.PartnerEntId = ""
	v.LastModDate = ""
	v.CrtDate = ""
	v.CrtIcName = ""
	v.Status = ""
	v.ModIcName = ""
	v.PartnerLevel = ""
	v.ModIcCode = ""
	v.PEntParId = ""
	v.CrtIcCode = ""
	v.ParRefEntId = ""
	v.LicenseTypeStr = ""
	v.AddrDetail = ""
	v.CountryName = ""
	v.AuditFlag = 0
	v.NetworkType = 0
	poolPEntParDto.Put(v)
}
