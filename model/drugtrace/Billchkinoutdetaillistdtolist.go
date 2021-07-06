package drugtrace

// Billchkinoutdetaillistdtolist 结构体
type Billchkinoutdetaillistdtolist struct {
	// 码列表
	CodeAndParentList []Codeandparentlist `json:"code_and_parent_list,omitempty" xml:"code_and_parent_list>codeandparentlist,omitempty"`
	// 有效期至
	ExpiredDate string `json:"expired_date,omitempty" xml:"expired_date,omitempty"`
	// 生产企业名称
	ProduceEntName string `json:"produce_ent_name,omitempty" xml:"produce_ent_name,omitempty"`
	// 子类编码
	ProdCode string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
	// 子类编码前7位
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 批次号
	ProductBatchNo string `json:"product_batch_no,omitempty" xml:"product_batch_no,omitempty"`
	// 药品id
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 药品名称
	PhysicName string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
	// 制剂单位
	PreparationsUnit string `json:"preparations_unit,omitempty" xml:"preparations_unit,omitempty"`
	// 包装规格
	TempPkgSpec string `json:"temp_pkg_spec,omitempty" xml:"temp_pkg_spec,omitempty"`
	// 最小制剂数量
	MinPreparationsCount string `json:"min_preparations_count,omitempty" xml:"min_preparations_count,omitempty"`
	// 最小包装数量
	MinPkgCount string `json:"min_pkg_count,omitempty" xml:"min_pkg_count,omitempty"`
	// 药品类型名称
	PhysicTypeName string `json:"physic_type_name,omitempty" xml:"physic_type_name,omitempty"`
	// 药品类型编码
	PhysicType string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
	// 国药准字
	ApproveNo string `json:"approve_no,omitempty" xml:"approve_no,omitempty"`
}
