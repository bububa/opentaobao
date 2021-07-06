package drugtrace

// SubType 结构体
type SubType struct {
	// 药品id
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 批准文号
	AuthorizedNo string `json:"authorized_no,omitempty" xml:"authorized_no,omitempty"`
	// 包装单位
	PackUnit string `json:"pack_unit,omitempty" xml:"pack_unit,omitempty"`
	// 包装规格
	PackageSpec string `json:"package_spec,omitempty" xml:"package_spec,omitempty"`
	// 药品详细类型
	PhysicDetailType string `json:"physic_detail_type,omitempty" xml:"physic_detail_type,omitempty"`
	// 有效期
	PhysicExpiry string `json:"physic_expiry,omitempty" xml:"physic_expiry,omitempty"`
	// 有效期单位
	PhysicExpiryUnit string `json:"physic_expiry_unit,omitempty" xml:"physic_expiry_unit,omitempty"`
	// 药品信息
	PhysicInfo string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
	// 药品类型
	PhysicType string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
	// 包装数量
	PkgNum string `json:"pkg_num,omitempty" xml:"pkg_num,omitempty"`
	// 制剂单位
	PrepnUnit string `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
	// 药品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 制剂规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 子类编码
	TypeNo string `json:"type_no,omitempty" xml:"type_no,omitempty"`
	// 资源码信息
	ResProdCodes *ResProdCodes `json:"res_prod_codes,omitempty" xml:"res_prod_codes,omitempty"`
}
