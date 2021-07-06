package film

// ActivityTagVo 结构体
type ActivityTagVo struct {
	// 活动内容
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 活动标：新促卡惠食星节
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
	// tagType
	TagType int64 `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
}
