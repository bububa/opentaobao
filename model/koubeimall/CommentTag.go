package koubeimall

import (
	"sync"
)

// CommentTag 结构体
type CommentTag struct {
	// 标签内容
	TagContent string `json:"tag_content,omitempty" xml:"tag_content,omitempty"`
	// 情感标示（1：正向 -1：负向）
	Emotion int64 `json:"emotion,omitempty" xml:"emotion,omitempty"`
	// 标签数量
	TagCount int64 `json:"tag_count,omitempty" xml:"tag_count,omitempty"`
}

var poolCommentTag = sync.Pool{
	New: func() any {
		return new(CommentTag)
	},
}

// GetCommentTag() 从对象池中获取CommentTag
func GetCommentTag() *CommentTag {
	return poolCommentTag.Get().(*CommentTag)
}

// ReleaseCommentTag 释放CommentTag
func ReleaseCommentTag(v *CommentTag) {
	v.TagContent = ""
	v.Emotion = 0
	v.TagCount = 0
	poolCommentTag.Put(v)
}
