package drugtrace

import (
	"sync"
)

// BaseInfoDto 结构体
type BaseInfoDto struct {
	// 药品信息
	PhysicInfo string `json:"physic_info,omitempty" xml:"physic_info,omitempty"`
	// 企业id
	RefEntId string `json:"ref_ent_id,omitempty" xml:"ref_ent_id,omitempty"`
	// 药品包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 药品制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 药品制剂类型
	PrepnType string `json:"prepn_type,omitempty" xml:"prepn_type,omitempty"`
	// 药品通用名称
	PhysicName string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
	// 药品包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
	// 药品有效期至
	ExprieDate string `json:"exprie_date,omitempty" xml:"exprie_date,omitempty"`
	// 药品生产批次号
	ProduceBatchNo string `json:"produce_batch_no,omitempty" xml:"produce_batch_no,omitempty"`
	// 药品生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 药品自类编码
	SubTypeNo string `json:"sub_type_no,omitempty" xml:"sub_type_no,omitempty"`
	// 药品编号
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 药品ID
	ProdId string `json:"prod_id,omitempty" xml:"prod_id,omitempty"`
	// 批准文号
	ApproveNo string `json:"approve_no,omitempty" xml:"approve_no,omitempty"`
	// 药品类型
	PhysicType string `json:"physic_type,omitempty" xml:"physic_type,omitempty"`
	// 原批准文号有效期
	ApprovalLicenceExpiryOld string `json:"approval_licence_expiry_old,omitempty" xml:"approval_licence_expiry_old,omitempty"`
	// 批准文号批准日期
	ApprovalLicenceDate string `json:"approval_licence_date,omitempty" xml:"approval_licence_date,omitempty"`
	// 批准文号有效期至
	ApprovalLicenceExpiry string `json:"approval_licence_expiry,omitempty" xml:"approval_licence_expiry,omitempty"`
	// 原批准文号
	ApprovalLicenceNoOld string `json:"approval_licence_no_old,omitempty" xml:"approval_licence_no_old,omitempty"`
}

var poolBaseInfoDto = sync.Pool{
	New: func() any {
		return new(BaseInfoDto)
	},
}

// GetBaseInfoDto() 从对象池中获取BaseInfoDto
func GetBaseInfoDto() *BaseInfoDto {
	return poolBaseInfoDto.Get().(*BaseInfoDto)
}

// ReleaseBaseInfoDto 释放BaseInfoDto
func ReleaseBaseInfoDto(v *BaseInfoDto) {
	v.PhysicInfo = ""
	v.RefEntId = ""
	v.PkgSpec = ""
	v.PrepnSpec = ""
	v.PrepnType = ""
	v.PhysicName = ""
	v.PkgRatio = ""
	v.ExprieDate = ""
	v.ProduceBatchNo = ""
	v.ProduceDate = ""
	v.SubTypeNo = ""
	v.ProductCode = ""
	v.ProdId = ""
	v.ApproveNo = ""
	v.PhysicType = ""
	v.ApprovalLicenceExpiryOld = ""
	v.ApprovalLicenceDate = ""
	v.ApprovalLicenceExpiry = ""
	v.ApprovalLicenceNoOld = ""
	poolBaseInfoDto.Put(v)
}
