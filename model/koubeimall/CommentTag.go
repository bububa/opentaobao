package koubeimall

// CommentTag 结构体
type CommentTag struct {
	// 情感标示（1：正向 -1：负向）
	Emotion int64 `json:"emotion,omitempty" xml:"emotion,omitempty"`
	// 标签数量
	TagCount int64 `json:"tag_count,omitempty" xml:"tag_count,omitempty"`
	// 标签内容
	TagContent string `json:"tag_content,omitempty" xml:"tag_content,omitempty"`
}
