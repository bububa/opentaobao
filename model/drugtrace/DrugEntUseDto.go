package drugtrace

import (
	"sync"
)

// DrugEntUseDto 结构体
type DrugEntUseDto struct {
	// 药品通用名称
	PhysicName string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
	// 药品英文名称
	PhysicNameEn string `json:"physic_name_en,omitempty" xml:"physic_name_en,omitempty"`
	// 药品商品名称
	ProdName string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	// 国家药品标识码
	CfdaDrugId string `json:"cfda_drug_id,omitempty" xml:"cfda_drug_id,omitempty"`
	// 药品本位码
	SdcCode string `json:"sdc_code,omitempty" xml:"sdc_code,omitempty"`
	// 剂型
	PrepnTypeDesc string `json:"prepn_type_desc,omitempty" xml:"prepn_type_desc,omitempty"`
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 药品有效期
	ExpiryTerm string `json:"expiry_term,omitempty" xml:"expiry_term,omitempty"`
	// 药品批准文号
	ApprovalLicenceNo string `json:"approval_licence_no,omitempty" xml:"approval_licence_no,omitempty"`
	// 药品批准文号有效期
	ApprovalLicenceExpiry string `json:"approval_licence_expiry,omitempty" xml:"approval_licence_expiry,omitempty"`
	// 药品生产日期
	ProductionDate string `json:"production_date,omitempty" xml:"production_date,omitempty"`
	// 药品有效期截止日期
	ExpiryDate string `json:"expiry_date,omitempty" xml:"expiry_date,omitempty"`
	// 药品生产批号
	ProductionBatch string `json:"production_batch,omitempty" xml:"production_batch,omitempty"`
	// 进口药品注册证号
	ImportRegCert string `json:"import_reg_cert,omitempty" xml:"import_reg_cert,omitempty"`
	// 进口药品注册证有效期
	ImportRegCertValidity string `json:"import_reg_cert_validity,omitempty" xml:"import_reg_cert_validity,omitempty"`
	// 进口药品批件号
	ImportAppCert string `json:"import_app_cert,omitempty" xml:"import_app_cert,omitempty"`
	// 进口药品批件有效期
	ImportAppCertValidity string `json:"import_app_cert_validity,omitempty" xml:"import_app_cert_validity,omitempty"`
	// S:已售出
	CodeStatus string `json:"code_status,omitempty" xml:"code_status,omitempty"`
	// 包装转换比
	PkgNum int64 `json:"pkg_num,omitempty" xml:"pkg_num,omitempty"`
	// 药品有效期
	ExprieLife int64 `json:"exprie_life,omitempty" xml:"exprie_life,omitempty"`
	// 药品有效期单位
	ExprieUnit int64 `json:"exprie_unit,omitempty" xml:"exprie_unit,omitempty"`
	// 药品注册分类
	DrugRegistrationClassfication int64 `json:"drug_registration_classfication,omitempty" xml:"drug_registration_classfication,omitempty"`
	// 国家基本药物标识
	NationalEssentialDrugsFlag int64 `json:"national_essential_drugs_flag,omitempty" xml:"national_essential_drugs_flag,omitempty"`
	// 特殊药品管理分类
	ControlledDrugManagementType int64 `json:"controlled_drug_management_type,omitempty" xml:"controlled_drug_management_type,omitempty"`
	// 处方药标识
	OtcFlag int64 `json:"otc_flag,omitempty" xml:"otc_flag,omitempty"`
	// 厂商信息
	AuthorizerEntInfo *BaseEntInfoDto `json:"authorizer_ent_info,omitempty" xml:"authorizer_ent_info,omitempty"`
	// 生产企业信息
	ProduceEntInfo *BaseEntInfoDto `json:"produce_ent_info,omitempty" xml:"produce_ent_info,omitempty"`
	// 分包装厂
	PackEntInfo *BaseEntInfoDto `json:"pack_ent_info,omitempty" xml:"pack_ent_info,omitempty"`
	// 代理企业
	AgentEntInfo *BaseEntInfoDto `json:"agent_ent_info,omitempty" xml:"agent_ent_info,omitempty"`
	// 零售企业信息
	RetailEntInfo *BaseEntInfoDto `json:"retail_ent_info,omitempty" xml:"retail_ent_info,omitempty"`
}

var poolDrugEntUseDto = sync.Pool{
	New: func() any {
		return new(DrugEntUseDto)
	},
}

// GetDrugEntUseDto() 从对象池中获取DrugEntUseDto
func GetDrugEntUseDto() *DrugEntUseDto {
	return poolDrugEntUseDto.Get().(*DrugEntUseDto)
}

// ReleaseDrugEntUseDto 释放DrugEntUseDto
func ReleaseDrugEntUseDto(v *DrugEntUseDto) {
	v.PhysicName = ""
	v.PhysicNameEn = ""
	v.ProdName = ""
	v.CfdaDrugId = ""
	v.SdcCode = ""
	v.PrepnTypeDesc = ""
	v.PrepnSpec = ""
	v.PkgSpec = ""
	v.ExpiryTerm = ""
	v.ApprovalLicenceNo = ""
	v.ApprovalLicenceExpiry = ""
	v.ProductionDate = ""
	v.ExpiryDate = ""
	v.ProductionBatch = ""
	v.ImportRegCert = ""
	v.ImportRegCertValidity = ""
	v.ImportAppCert = ""
	v.ImportAppCertValidity = ""
	v.CodeStatus = ""
	v.PkgNum = 0
	v.ExprieLife = 0
	v.ExprieUnit = 0
	v.DrugRegistrationClassfication = 0
	v.NationalEssentialDrugsFlag = 0
	v.ControlledDrugManagementType = 0
	v.OtcFlag = 0
	v.AuthorizerEntInfo = nil
	v.ProduceEntInfo = nil
	v.PackEntInfo = nil
	v.AgentEntInfo = nil
	v.RetailEntInfo = nil
	poolDrugEntUseDto.Put(v)
}
