package drugtrace

// AlibabaalihealthdrugkytlistauthsResult 结构体
type AlibabaalihealthdrugkytlistauthsResult struct {
	// 企业ID
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 企业名称
	EntCapitalName string `json:"ent_capital_name,omitempty" xml:"ent_capital_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 地市编码
	DictRegionCode string `json:"dict_region_code,omitempty" xml:"dict_region_code,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 省
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 注册编码
	RegRegionCode string `json:"reg_region_code,omitempty" xml:"reg_region_code,omitempty"`
	// 角色类型
	UserRoleType string `json:"user_role_type,omitempty" xml:"user_role_type,omitempty"`
	// 是否入网
	IsNetwork string `json:"is_network,omitempty" xml:"is_network,omitempty"`
	// 企业唯一标识
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
}
