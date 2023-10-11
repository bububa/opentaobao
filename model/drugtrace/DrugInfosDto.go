package drugtrace

// DrugInfosDto 结构体
type DrugInfosDto struct {
	// 码信息
	CodeInfoListDtoList []CodeInfoListDto `json:"code_info_list_dto_list,omitempty" xml:"code_info_list_dto_list>code_info_list_dto,omitempty"`
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 生产企业名称
	ProductEntName string `json:"product_ent_name,omitempty" xml:"product_ent_name,omitempty"`
	// 产品包装规格
	PackageSpec string `json:"package_spec,omitempty" xml:"package_spec,omitempty"`
	// 药品商品名
	ProdName string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 制剂单位编码
	PrepnUnit string `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
	// 批次号
	ProduceBatchNo string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
	// 药品标识
	ProdSeqNo string `json:"prod_seq_no,omitempty" xml:"prod_seq_no,omitempty"`
	// 药品标识
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 有效期至
	ValidEndDate string `json:"valid_end_date,omitempty" xml:"valid_end_date,omitempty"`
	// 按最小包装单位统计数量
	LeastPkgAmount string `json:"least_pkg_amount,omitempty" xml:"least_pkg_amount,omitempty"`
	// 按最小制剂单位统计数量
	LeastPrepnAmount string `json:"least_prepn_amount,omitempty" xml:"least_prepn_amount,omitempty"`
	// 批准文号
	ApprovalNo string `json:"approval_no,omitempty" xml:"approval_no,omitempty"`
}
