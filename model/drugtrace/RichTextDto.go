package drugtrace

import (
	"sync"
)

// RichTextDto 结构体
type RichTextDto struct {
	// 图片
	Pictures []string `json:"pictures,omitempty" xml:"pictures>string,omitempty"`
	// 文字
	Text string `json:"text,omitempty" xml:"text,omitempty"`
}

var poolRichTextDto = sync.Pool{
	New: func() any {
		return new(RichTextDto)
	},
}

// GetRichTextDto() 从对象池中获取RichTextDto
func GetRichTextDto() *RichTextDto {
	return poolRichTextDto.Get().(*RichTextDto)
}

// ReleaseRichTextDto 释放RichTextDto
func ReleaseRichTextDto(v *RichTextDto) {
	v.Pictures = v.Pictures[:0]
	v.Text = ""
	poolRichTextDto.Put(v)
}
