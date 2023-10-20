package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytDrDrugrecalModel 结构体
type AlibabaAlihealthDrugKytDrDrugrecalModel struct {
	// 药品剂型
	PrepnTypeDesc string `json:"prepn_type_desc,omitempty" xml:"prepn_type_desc,omitempty"`
	// 召回结束时间
	RecallEndDate string `json:"recall_end_date,omitempty" xml:"recall_end_date,omitempty"`
	// 药品批次编号
	ProductBatchNos string `json:"product_batch_nos,omitempty" xml:"product_batch_nos,omitempty"`
	// 药品生产企业名称
	ProdcutEntName string `json:"prodcut_ent_name,omitempty" xml:"prodcut_ent_name,omitempty"`
	// 药品批准文号
	ApproveNo string `json:"approve_no,omitempty" xml:"approve_no,omitempty"`
	// 药品通用名称
	PhysicName string `json:"physic_name,omitempty" xml:"physic_name,omitempty"`
	// 召回范围(0全国1省2市)
	RecallRange string `json:"recall_range,omitempty" xml:"recall_range,omitempty"`
	// 是否向消费者公开(1公开0不公开)
	RecallOpen string `json:"recall_open,omitempty" xml:"recall_open,omitempty"`
	// 批准状态(0未批准1已批准)
	RecallAuditStatus string `json:"recall_audit_status,omitempty" xml:"recall_audit_status,omitempty"`
	// 包装规格
	PkgSpec string `json:"pkg_spec,omitempty" xml:"pkg_spec,omitempty"`
	// 制剂规格
	PrepnSpec string `json:"prepn_spec,omitempty" xml:"prepn_spec,omitempty"`
	// 召回ID
	RecallInfoId string `json:"recall_info_id,omitempty" xml:"recall_info_id,omitempty"`
	// 召回企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 召回性质(1实施召回2模拟召回)
	RecallNature string `json:"recall_nature,omitempty" xml:"recall_nature,omitempty"`
	// 召回原因
	RecallReason string `json:"recall_reason,omitempty" xml:"recall_reason,omitempty"`
	// 召回单类型(1主动召回2责令召回)
	RecallBillType string `json:"recall_bill_type,omitempty" xml:"recall_bill_type,omitempty"`
	// 召回状态(0未召回1召回中2召回结束)
	RecallStatus string `json:"recall_status,omitempty" xml:"recall_status,omitempty"`
	// 召回级别(1一级2二级3三级)
	RecallLevel string `json:"recall_level,omitempty" xml:"recall_level,omitempty"`
}

var poolAlibabaAlihealthDrugKytDrDrugrecalModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytDrDrugrecalModel)
	},
}

// GetAlibabaAlihealthDrugKytDrDrugrecalModel() 从对象池中获取AlibabaAlihealthDrugKytDrDrugrecalModel
func GetAlibabaAlihealthDrugKytDrDrugrecalModel() *AlibabaAlihealthDrugKytDrDrugrecalModel {
	return poolAlibabaAlihealthDrugKytDrDrugrecalModel.Get().(*AlibabaAlihealthDrugKytDrDrugrecalModel)
}

// ReleaseAlibabaAlihealthDrugKytDrDrugrecalModel 释放AlibabaAlihealthDrugKytDrDrugrecalModel
func ReleaseAlibabaAlihealthDrugKytDrDrugrecalModel(v *AlibabaAlihealthDrugKytDrDrugrecalModel) {
	v.PrepnTypeDesc = ""
	v.RecallEndDate = ""
	v.ProductBatchNos = ""
	v.ProdcutEntName = ""
	v.ApproveNo = ""
	v.PhysicName = ""
	v.RecallRange = ""
	v.RecallOpen = ""
	v.RecallAuditStatus = ""
	v.PkgSpec = ""
	v.PrepnSpec = ""
	v.RecallInfoId = ""
	v.EntName = ""
	v.RecallNature = ""
	v.RecallReason = ""
	v.RecallBillType = ""
	v.RecallStatus = ""
	v.RecallLevel = ""
	poolAlibabaAlihealthDrugKytDrDrugrecalModel.Put(v)
}
