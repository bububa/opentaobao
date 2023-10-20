package legalsuit

import (
	"sync"
)

// DominationModel 结构体
type DominationModel struct {
	// 管辖裁决书附件
	DominationDissentRulingList []FileModel `json:"domination_dissent_ruling_list,omitempty" xml:"domination_dissent_ruling_list>file_model,omitempty"`
	// 管辖文书列表
	AttachmentList []FileModel `json:"attachment_list,omitempty" xml:"attachment_list>file_model,omitempty"`
	// 管辖裁决书数量
	DominationissentRulingCount string `json:"dominationissent_ruling_count,omitempty" xml:"dominationissent_ruling_count,omitempty"`
	// 管辖文书
	AttachmentCount string `json:"attachment_count,omitempty" xml:"attachment_count,omitempty"`
	// 备注
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 快递编号
	ExpressNum string `json:"express_num,omitempty" xml:"express_num,omitempty"`
	// 快递公司
	ExpressCompany string `json:"express_company,omitempty" xml:"express_company,omitempty"`
	// 法官
	Judge string `json:"judge,omitempty" xml:"judge,omitempty"`
	// 是否开庭
	IsCourt string `json:"is_court,omitempty" xml:"is_court,omitempty"`
	// 供应商code
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 案件ID
	SuitId int64 `json:"suit_id,omitempty" xml:"suit_id,omitempty"`
	// 委托ID
	EntrustId int64 `json:"entrust_id,omitempty" xml:"entrust_id,omitempty"`
	// 管辖ID
	DominationId int64 `json:"domination_id,omitempty" xml:"domination_id,omitempty"`
}

var poolDominationModel = sync.Pool{
	New: func() any {
		return new(DominationModel)
	},
}

// GetDominationModel() 从对象池中获取DominationModel
func GetDominationModel() *DominationModel {
	return poolDominationModel.Get().(*DominationModel)
}

// ReleaseDominationModel 释放DominationModel
func ReleaseDominationModel(v *DominationModel) {
	v.DominationDissentRulingList = v.DominationDissentRulingList[:0]
	v.AttachmentList = v.AttachmentList[:0]
	v.DominationissentRulingCount = ""
	v.AttachmentCount = ""
	v.Description = ""
	v.ExpressNum = ""
	v.ExpressCompany = ""
	v.Judge = ""
	v.IsCourt = ""
	v.SupplierCode = ""
	v.SuitId = 0
	v.EntrustId = 0
	v.DominationId = 0
	poolDominationModel.Put(v)
}
