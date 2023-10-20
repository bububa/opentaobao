package tmallgenie

import (
	"sync"
)

// Note 结构体
type Note struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// memo状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 记事本内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 记事本主题
	Topic string `json:"topic,omitempty" xml:"topic,omitempty"`
	// memo_ID
	MemoId int64 `json:"memo_id,omitempty" xml:"memo_id,omitempty"`
}

var poolNote = sync.Pool{
	New: func() any {
		return new(Note)
	},
}

// GetNote() 从对象池中获取Note
func GetNote() *Note {
	return poolNote.Get().(*Note)
}

// ReleaseNote 释放Note
func ReleaseNote(v *Note) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Uuid = ""
	v.Status = ""
	v.Content = ""
	v.Topic = ""
	v.MemoId = 0
	poolNote.Put(v)
}
