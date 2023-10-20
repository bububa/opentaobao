package legalsuit

import (
	"sync"
)

// AfterCourtDetailModel 结构体
type AfterCourtDetailModel struct {
	// 附件
	NonEviAttachmentList []FileModel `json:"non_evi_attachment_list,omitempty" xml:"non_evi_attachment_list>file_model,omitempty"`
	// 附件
	EviAttachmentList []FileModel `json:"evi_attachment_list,omitempty" xml:"evi_attachment_list>file_model,omitempty"`
	// 提交时间
	SubmitDate string `json:"submit_date,omitempty" xml:"submit_date,omitempty"`
	// 质证情况
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 证明目的
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 材料简介
	Intro string `json:"intro,omitempty" xml:"intro,omitempty"`
	// 材料名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 附件数量
	NonEviAttachmentCount int64 `json:"non_evi_attachment_count,omitempty" xml:"non_evi_attachment_count,omitempty"`
	// 附件数量
	EviAttachmentCount int64 `json:"evi_attachment_count,omitempty" xml:"evi_attachment_count,omitempty"`
}

var poolAfterCourtDetailModel = sync.Pool{
	New: func() any {
		return new(AfterCourtDetailModel)
	},
}

// GetAfterCourtDetailModel() 从对象池中获取AfterCourtDetailModel
func GetAfterCourtDetailModel() *AfterCourtDetailModel {
	return poolAfterCourtDetailModel.Get().(*AfterCourtDetailModel)
}

// ReleaseAfterCourtDetailModel 释放AfterCourtDetailModel
func ReleaseAfterCourtDetailModel(v *AfterCourtDetailModel) {
	v.NonEviAttachmentList = v.NonEviAttachmentList[:0]
	v.EviAttachmentList = v.EviAttachmentList[:0]
	v.SubmitDate = ""
	v.Description = ""
	v.Content = ""
	v.Intro = ""
	v.Name = ""
	v.NonEviAttachmentCount = 0
	v.EviAttachmentCount = 0
	poolAfterCourtDetailModel.Put(v)
}
