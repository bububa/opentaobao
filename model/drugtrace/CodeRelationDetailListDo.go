package drugtrace

// CodeRelationDetailListDo 结构体
type CodeRelationDetailListDo struct {
	// 包装规格= 数量+ 制剂单位/包装单位
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 制剂数量
	BatchNum string `json:"batch_num,omitempty" xml:"batch_num,omitempty"`
	// 生产负责人
	DutyMan string `json:"duty_man,omitempty" xml:"duty_man,omitempty"`
	// 激活数量(激活总量)
	ActiveCount string `json:"active_count,omitempty" xml:"active_count,omitempty"`
	// 最大包装数量
	OtherNum string `json:"other_num,omitempty" xml:"other_num,omitempty"`
	// 小码数量
	SmallNum string `json:"small_num,omitempty" xml:"small_num,omitempty"`
	// 生产线
	ProduceLine string `json:"produce_line,omitempty" xml:"produce_line,omitempty"`
	// 生产车间
	ProduceFactory string `json:"produce_factory,omitempty" xml:"produce_factory,omitempty"`
	// 包装比例信息
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 有效期至
	ExprieDate string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
	// 生产批次号
	ProduceBatchNo string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
	// 药品商品基本信息标识
	DrugId string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
	// 关联关系文件信息标识
	CodeActiveInfoId string `json:"code_active_info_id,omitempty" xml:"code_active_info_id,omitempty"`
	// 药品激活信息标识
	CodeActiveInfoListId string `json:"code_active_info_list_id,omitempty" xml:"code_active_info_list_id,omitempty"`
	// 包装单位
	PkgUnitDesc string `json:"pkg_unit_desc,omitempty" xml:"pkg_unit_desc,omitempty"`
	// 制剂单位
	PrepnUnitDesc string `json:"prepn_unit_desc,omitempty" xml:"prepn_unit_desc,omitempty"`
}
