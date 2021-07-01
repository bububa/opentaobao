package iot

// Like 结构体
type Like struct {
	// 收藏音频源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 收藏音频id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 收藏音频专辑名
	Album string `json:"album,omitempty" xml:"album,omitempty"`
	// 收藏音频演唱者
	Author string `json:"author,omitempty" xml:"author,omitempty"`
	// 收藏音频名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
