package drugtrace

// SubTypeList 结构体
type SubTypeList struct {
	// 码列表
	CodeResList []CodeResList `json:"code_res_list,omitempty" xml:"code_res_list>code_res_list,omitempty"`
	// 制剂单位
	PrepnUnit string `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
	// 包装规格
	PackageSpec string `json:"package_spec,omitempty" xml:"package_spec,omitempty"`
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 企业药品ID
	ProdSeqNo string `json:"prod_seq_no,omitempty" xml:"prod_seq_no,omitempty"`
	// 批准文号
	ApproveNo string `json:"approve_no,omitempty" xml:"approve_no,omitempty"`
	// 药品详情类型
	PhysicDetailType string `json:"physic_detail_type,omitempty" xml:"physic_detail_type,omitempty"`
	// 包装单位
	PackUnit string `json:"pack_unit,omitempty" xml:"pack_unit,omitempty"`
	// 药品ID
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 包装单位
	PackUnitName string `json:"pack_unit_name,omitempty" xml:"pack_unit_name,omitempty"`
	// 制剂描述
	PrepnDesc string `json:"prepn_desc,omitempty" xml:"prepn_desc,omitempty"`
	// 制剂单位描述
	PrepnUnitName string `json:"prepn_unit_name,omitempty" xml:"prepn_unit_name,omitempty"`
	// 子类型
	SubTypeNo string `json:"sub_type_no,omitempty" xml:"sub_type_no,omitempty"`
}
