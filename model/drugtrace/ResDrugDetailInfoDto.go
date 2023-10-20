package drugtrace

import (
	"sync"
)

// ResDrugDetailInfoDto 结构体
type ResDrugDetailInfoDto struct {
	// 包装单位描述
	PkgUnitDesc string `json:"pkg_unit_desc,omitempty" xml:"pkg_unit_desc,omitempty"`
	// 批准文号有效期至
	ApprovalLicenceExpiry string `json:"approval_licence_expiry,omitempty" xml:"approval_licence_expiry,omitempty"`
	// 有效期单位(详见码表) 1:日，2：月，3：年
	ExprieUnit string `json:"exprie_unit,omitempty" xml:"exprie_unit,omitempty"`
	// 药品详细类型
	PhysicDetailType string `json:"physic_detail_type,omitempty" xml:"physic_detail_type,omitempty"`
	// 包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 药品基本信息id
	DrugBaseInfoId string `json:"drug_base_info_id,omitempty" xml:"drug_base_info_id,omitempty"`
	// 状态 1正常,0停用
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 企业唯一标识
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 药品类型描叙（普通药品）
	PhysicTypeDesc string `json:"physic_type_desc,omitempty" xml:"physic_type_desc,omitempty"`
	// 包装数量
	PkgNum string `json:"pkg_num,omitempty" xml:"pkg_num,omitempty"`
	// 创建时间(为企业提供数据)
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 企业ID
	EntId string `json:"ent_id,omitempty" xml:"ent_id,omitempty"`
	// 制剂单位类型描述
	PrepnUnitDesc string `json:"prepn_unit_desc,omitempty" xml:"prepn_unit_desc,omitempty"`
	// 年赋码量
	AnnCodeAmt string `json:"ann_code_amt,omitempty" xml:"ann_code_amt,omitempty"`
	// 药品批准文号ID(药品目录)
	DrugApprovalInfoId string `json:"drug_approval_info_id,omitempty" xml:"drug_approval_info_id,omitempty"`
	// 药品本位码
	SdcCode string `json:"sdc_code,omitempty" xml:"sdc_code,omitempty"`
	// 药品有效期数(月/年)
	ExprieLife string `json:"exprie_life,omitempty" xml:"exprie_life,omitempty"`
	// 商品名称
	ProdName string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	// 药品类型(详见码表) 1：特殊药品原料药，2：特殊药品制剂，3：普通药品，9：未分类
	PhysicType string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
	// 生产企业详细地址
	RegRegionDetail string `json:"reg_region_detail,omitempty" xml:"reg_region_detail,omitempty"`
	// 申请码标识:1未申请,2已申请
	PhysicNatCode string `json:"physic_nat_code,omitempty" xml:"physic_nat_code,omitempty"`
	// 药品ID
	DrugEntBaseInfoId string `json:"drug_ent_base_info_id,omitempty" xml:"drug_ent_base_info_id,omitempty"`
	// 制剂单位类型(详见码表)
	PrepnUnit string `json:"prepn_unit,omitempty" xml:"prepn_unit,omitempty"`
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 剂型(制剂类型)
	PrepnType string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
	// 企业药品ID
	DrugEntBaseId string `json:"drug_ent_base_id,omitempty" xml:"drug_ent_base_id,omitempty"`
	// 剂型描述
	PrepnTypeDesc string `json:"prepn_type_desc,omitempty" xml:"prepn_type_desc,omitempty"`
	// 药品通用名称
	PhysicName string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
	// 包装单位
	PkgUnit string `json:"pkg_unit,omitempty" xml:"pkg_unit,omitempty"`
	// 包装规格常规
	PkgSpecCrit string `json:"pkg_spec_crit,omitempty" xml:"pkg_spec_crit,omitempty"`
	// 是否授权0未授权,1已授权
	AuthorizerFlag string `json:"authorizer_flag,omitempty" xml:"authorizer_flag,omitempty"`
	// 修改时间
	ModDate string `json:"mod_date,omitempty" xml:"mod_date,omitempty"`
	// 批准文号
	ApprovalLicenceNo string `json:"approval_licence_no,omitempty" xml:"approval_licence_no,omitempty"`
	// 药品信息
	PhysicInfo string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
	// 药品自类编码
	ProdCode string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
	// 批准文号类型
	ApprovalLicenceType string `json:"approval_licence_type,omitempty" xml:"approval_licence_type,omitempty"`
}

var poolResDrugDetailInfoDto = sync.Pool{
	New: func() any {
		return new(ResDrugDetailInfoDto)
	},
}

// GetResDrugDetailInfoDto() 从对象池中获取ResDrugDetailInfoDto
func GetResDrugDetailInfoDto() *ResDrugDetailInfoDto {
	return poolResDrugDetailInfoDto.Get().(*ResDrugDetailInfoDto)
}

// ReleaseResDrugDetailInfoDto 释放ResDrugDetailInfoDto
func ReleaseResDrugDetailInfoDto(v *ResDrugDetailInfoDto) {
	v.PkgUnitDesc = ""
	v.ApprovalLicenceExpiry = ""
	v.ExprieUnit = ""
	v.PhysicDetailType = ""
	v.PkgSpec = ""
	v.EntName = ""
	v.DrugBaseInfoId = ""
	v.Status = ""
	v.RefEntId = ""
	v.PhysicTypeDesc = ""
	v.PkgNum = ""
	v.CrtDate = ""
	v.EntId = ""
	v.PrepnUnitDesc = ""
	v.AnnCodeAmt = ""
	v.DrugApprovalInfoId = ""
	v.SdcCode = ""
	v.ExprieLife = ""
	v.ProdName = ""
	v.PhysicType = ""
	v.RegRegionDetail = ""
	v.PhysicNatCode = ""
	v.DrugEntBaseInfoId = ""
	v.PrepnUnit = ""
	v.PrepnSpec = ""
	v.PrepnType = ""
	v.DrugEntBaseId = ""
	v.PrepnTypeDesc = ""
	v.PhysicName = ""
	v.PkgUnit = ""
	v.PkgSpecCrit = ""
	v.AuthorizerFlag = ""
	v.ModDate = ""
	v.ApprovalLicenceNo = ""
	v.PhysicInfo = ""
	v.ProdCode = ""
	v.ApprovalLicenceType = ""
	poolResDrugDetailInfoDto.Put(v)
}
