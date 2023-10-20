package legalcase

import (
	"sync"
)

// NoticeModel 结构体
type NoticeModel struct {
	// 附件
	Attachment []FileModel `json:"attachment,omitempty" xml:"attachment>file_model,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 消息唯一id
	NoticeId string `json:"notice_id,omitempty" xml:"notice_id,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolNoticeModel = sync.Pool{
	New: func() any {
		return new(NoticeModel)
	},
}

// GetNoticeModel() 从对象池中获取NoticeModel
func GetNoticeModel() *NoticeModel {
	return poolNoticeModel.Get().(*NoticeModel)
}

// ReleaseNoticeModel 释放NoticeModel
func ReleaseNoticeModel(v *NoticeModel) {
	v.Attachment = v.Attachment[:0]
	v.Creator = ""
	v.NoticeId = ""
	v.Content = ""
	poolNoticeModel.Put(v)
}
