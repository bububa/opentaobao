package legalcase

// NoticeModel 结构体
type NoticeModel struct {
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 附件
	Attachment []FileModel `json:"attachment,omitempty" xml:"attachment>file_model,omitempty"`
	// 消息唯一id
	NoticeId string `json:"notice_id,omitempty" xml:"notice_id,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}
