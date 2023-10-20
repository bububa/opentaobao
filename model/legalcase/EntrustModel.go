package legalcase

import (
	"sync"
)

// EntrustModel 结构体
type EntrustModel struct {
	// 附件
	AttachmentList []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
	// 案件ids
	CaseIds []int64 `json:"case_ids,omitempty" xml:"case_ids>int64,omitempty"`
	// 委托项
	EntrustTypes []string `json:"entrust_types,omitempty" xml:"entrust_types>string,omitempty"`
	// 调解截止时间
	Deadline string `json:"deadline,omitempty" xml:"deadline,omitempty"`
	// 备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 委托编号
	EntrustCode string `json:"entrust_code,omitempty" xml:"entrust_code,omitempty"`
	// 委托方名称
	EntrustParty string `json:"entrust_party,omitempty" xml:"entrust_party,omitempty"`
	// 委托人
	EntrustPeople string `json:"entrust_people,omitempty" xml:"entrust_people,omitempty"`
	// 委托时间
	EntrustTime string `json:"entrust_time,omitempty" xml:"entrust_time,omitempty"`
	// lvms委托编号
	LvmsEntrustCode string `json:"lvms_entrust_code,omitempty" xml:"lvms_entrust_code,omitempty"`
	// 主要负责律师
	MainLawyer string `json:"main_lawyer,omitempty" xml:"main_lawyer,omitempty"`
	// 策略和建议
	Suggest string `json:"suggest,omitempty" xml:"suggest,omitempty"`
	// 供应商编号
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
}

var poolEntrustModel = sync.Pool{
	New: func() any {
		return new(EntrustModel)
	},
}

// GetEntrustModel() 从对象池中获取EntrustModel
func GetEntrustModel() *EntrustModel {
	return poolEntrustModel.Get().(*EntrustModel)
}

// ReleaseEntrustModel 释放EntrustModel
func ReleaseEntrustModel(v *EntrustModel) {
	v.AttachmentList = v.AttachmentList[:0]
	v.CaseIds = v.CaseIds[:0]
	v.EntrustTypes = v.EntrustTypes[:0]
	v.Deadline = ""
	v.Description = ""
	v.EntrustCode = ""
	v.EntrustParty = ""
	v.EntrustPeople = ""
	v.EntrustTime = ""
	v.LvmsEntrustCode = ""
	v.MainLawyer = ""
	v.Suggest = ""
	v.SupplierCode = ""
	v.SupplierName = ""
	poolEntrustModel.Put(v)
}
