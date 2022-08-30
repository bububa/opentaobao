package drugtrace

// DrugTableDto 结构体
type DrugTableDto struct {
	// 子列表
	SubTypeList []SubTypeList `json:"sub_type_list,omitempty" xml:"sub_type_list>sub_type_list,omitempty"`
	// 制剂类型描述
	PrepnTypeDesc string `json:"prepn_type_desc,omitempty" xml:"prepn_type_desc,omitempty"`
	// 药品类型描述
	PhysicTypeDesc string `json:"physic_type_desc,omitempty" xml:"physic_type_desc,omitempty"`
	// 药品名称
	PhysicName string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
	// 药品自类编码
	ProdCode string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
	// 企业主键
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 商品名称
	ProdName string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	// 修改日期
	ModDate string `json:"mod_date,omitempty" xml:"mod_date,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 包装单位描述
	PkgUnitDesc string `json:"pkg_unit_desc,omitempty" xml:"pkg_unit_desc,omitempty"`
	// 药品类型详情描述
	PhysicDetailTypeDesc string `json:"physic_detail_type_desc,omitempty" xml:"physic_detail_type_desc,omitempty"`
	// 制剂单位描述
	PrepnUnitDesc string `json:"prepn_unit_desc,omitempty" xml:"prepn_unit_desc,omitempty"`
	// 药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类
	PhysicType int64 `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
	// 药品详细类型
	PhysicDetailType int64 `json:"physic_detail_type,omitempty" xml:"physic_detail_type,omitempty"`
}
