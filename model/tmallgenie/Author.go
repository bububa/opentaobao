package tmallgenie

import (
	"sync"
)

// Author 结构体
type Author struct {
	// 内容作者，对应音乐为作词、作曲人，对应小说故事为原著作者
	ContentAuthor string `json:"content_author,omitempty" xml:"content_author,omitempty"`
	// 主播，演唱者，演播者
	VoiceAuthor string `json:"voice_author,omitempty" xml:"voice_author,omitempty"`
}

var poolAuthor = sync.Pool{
	New: func() any {
		return new(Author)
	},
}

// GetAuthor() 从对象池中获取Author
func GetAuthor() *Author {
	return poolAuthor.Get().(*Author)
}

// ReleaseAuthor 释放Author
func ReleaseAuthor(v *Author) {
	v.ContentAuthor = ""
	v.VoiceAuthor = ""
	poolAuthor.Put(v)
}
