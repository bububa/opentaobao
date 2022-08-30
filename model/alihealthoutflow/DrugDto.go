package alihealthoutflow

// DrugDto 结构体
type DrugDto struct {
	// 用法用量
	DrugUsageList []DrugUsageVo `json:"drug_usage_list,omitempty" xml:"drug_usage_list>drug_usage_vo,omitempty"`
	// 规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 总量
	Total string `json:"total,omitempty" xml:"total,omitempty"`
	// 药品通用名
	CommonDrugName string `json:"common_drug_name,omitempty" xml:"common_drug_name,omitempty"`
	// 药品名
	DrugName string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
	// 剂型
	DoseForm string `json:"dose_form,omitempty" xml:"dose_form,omitempty"`
	// 药品通用名拼音
	DrugCommonNamePy string `json:"drug_common_name_py,omitempty" xml:"drug_common_name_py,omitempty"`
	// 药品名拼音
	DrugNamePy string `json:"drug_name_py,omitempty" xml:"drug_name_py,omitempty"`
	// 省平台ID
	ProvinceDrugCode string `json:"province_drug_code,omitempty" xml:"province_drug_code,omitempty"`
	// 适应症
	Indication string `json:"indication,omitempty" xml:"indication,omitempty"`
	// HIS药品名称
	HisDrugName string `json:"his_drug_name,omitempty" xml:"his_drug_name,omitempty"`
	// 每次剂量单位
	SingleDosageUnit string `json:"single_dosage_unit,omitempty" xml:"single_dosage_unit,omitempty"`
	// 每次剂量值
	SingleDosage string `json:"single_dosage,omitempty" xml:"single_dosage,omitempty"`
	// 包装规格--最小销售单位
	PackSpecSaleUnit string `json:"pack_spec_sale_unit,omitempty" xml:"pack_spec_sale_unit,omitempty"`
	// 包装规格--最小使用单位
	PackSpecUseUnit string `json:"pack_spec_use_unit,omitempty" xml:"pack_spec_use_unit,omitempty"`
	// 包装规格--包装数量
	PackSpecAmount string `json:"pack_spec_amount,omitempty" xml:"pack_spec_amount,omitempty"`
	// 包装单位-装量单位
	PackUnit string `json:"pack_unit,omitempty" xml:"pack_unit,omitempty"`
	// 包装单位-装量值
	PackQty string `json:"pack_qty,omitempty" xml:"pack_qty,omitempty"`
	// 单价
	RetailPrice string `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 医保编号
	MedicareNo string `json:"medicare_no,omitempty" xml:"medicare_no,omitempty"`
	// 用药途径编码
	DrugUsageCode string `json:"drug_usage_code,omitempty" xml:"drug_usage_code,omitempty"`
	// 用药途径
	DrugUsage string `json:"drug_usage,omitempty" xml:"drug_usage,omitempty"`
	// 频次编码
	FrequencyCode string `json:"frequency_code,omitempty" xml:"frequency_code,omitempty"`
	// 频次
	Frequency string `json:"frequency,omitempty" xml:"frequency,omitempty"`
	// 药品本位码
	DrugBasicCode string `json:"drug_basic_code,omitempty" xml:"drug_basic_code,omitempty"`
	// 剂型
	DosageForm string `json:"dosage_form,omitempty" xml:"dosage_form,omitempty"`
	// 国药准字
	ApprovalNumber string `json:"approval_number,omitempty" xml:"approval_number,omitempty"`
	// 生产企业
	Manufacture string `json:"manufacture,omitempty" xml:"manufacture,omitempty"`
	// 处方展示规格
	PackSpec string `json:"pack_spec,omitempty" xml:"pack_spec,omitempty"`
	// 医院编码
	DrugCode string `json:"drug_code,omitempty" xml:"drug_code,omitempty"`
	// 药品通用名
	DrugCommonName string `json:"drug_common_name,omitempty" xml:"drug_common_name,omitempty"`
	// 药品分类码
	DrugClassCode string `json:"drug_class_code,omitempty" xml:"drug_class_code,omitempty"`
	// 用药单位
	DoseUnit string `json:"dose_unit,omitempty" xml:"dose_unit,omitempty"`
	// 用药数量
	Dose string `json:"dose,omitempty" xml:"dose,omitempty"`
	// 用法用量+医嘱
	DoseUsageAdvice string `json:"dose_usage_advice,omitempty" xml:"dose_usage_advice,omitempty"`
	// 药品厂家
	Manufactures string `json:"manufactures,omitempty" xml:"manufactures,omitempty"`
	// 商品名称
	ProdName string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	// 药品数量（必选）
	Count string `json:"count,omitempty" xml:"count,omitempty"`
	// 核销药品单价
	Price *BigDecimal `json:"price,omitempty" xml:"price,omitempty"`
	// 是否允许电子处方
	ElectricPres bool `json:"electric_pres,omitempty" xml:"electric_pres,omitempty"`
}
