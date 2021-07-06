package wenyuvideo

// PersonDo 结构体
type PersonDo struct {
	// 人物名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图片
	ThumbUrl string `json:"thumb_url,omitempty" xml:"thumb_url,omitempty"`
	// 300*300人物头像
	ThumbUrlLg string `json:"thumb_url_lg,omitempty" xml:"thumb_url_lg,omitempty"`
	// 人物海报
	PosterUrl string `json:"poster_url,omitempty" xml:"poster_url,omitempty"`
	// 人物写真
	PosterUrlH string `json:"poster_url_h,omitempty" xml:"poster_url_h,omitempty"`
	// 简介
	PersonDesc string `json:"person_desc,omitempty" xml:"person_desc,omitempty"`
	// 人物ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
