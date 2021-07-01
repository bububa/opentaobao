package drugtrace

// ProductInfoList 结构体
type ProductInfoList struct {
	// 药品详细类型
	PhysicDetailType int64 `json:"physic_detail_type,omitempty" xml:"physic_detail_type,omitempty"`
	// 69码
	Ean string `json:"ean,omitempty" xml:"ean,omitempty"`
	// 国药准字
	AuthorizedNo string `json:"authorized_no,omitempty" xml:"authorized_no,omitempty"`
	// 企业id
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 药品类型
	PhysicType string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
	// 包装数量
	PkgNum int64 `json:"pkg_num,omitempty" xml:"pkg_num,omitempty"`
	// 制剂单位
	PrepnUnit string `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
	// 商品码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 药品编号
	PhysicCode string `json:"physic_code,omitempty" xml:"physic_code,omitempty"`
	// 有效期至
	ExprieDate string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
	// 生产日期
	ProductDate string `json:"product_date,omitempty" xml:"product_date,omitempty"`
	// 批次号
	ProductBatchNo string `json:"product_batch_no,omitempty" xml:"product_batch_no,omitempty"`
	// 药品数量
	DrugNum int64 `json:"drug_num,omitempty" xml:"drug_num,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 包装规格单位
	PkgSpecUnit string `json:"pkg_spec_unit,omitempty" xml:"pkg_spec_unit,omitempty"`
	// 包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 制剂类型
	PrepnType string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
	// 药品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 药品通用名
	DrugPhysicName string `json:"drug_physic_name,omitempty" xml:"drug_physic_name,omitempty"`
	// 生产信息id
	ProduceInfoId string `json:"produce_info_id,omitempty" xml:"produce_info_id,omitempty"`
	// 药品id
	DrugId string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
}
