package nlp

import (
	"sync"
)

// Text 结构体
type Text struct {
	// 业务处理ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 文本内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 文本类型1-评论 2-订单留言 9-其他
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolText = sync.Pool{
	New: func() any {
		return new(Text)
	},
}

// GetText() 从对象池中获取Text
func GetText() *Text {
	return poolText.Get().(*Text)
}

// ReleaseText 释放Text
func ReleaseText(v *Text) {
	v.Id = ""
	v.Content = ""
	v.Type = 0
	poolText.Put(v)
}
