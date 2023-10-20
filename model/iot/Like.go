package iot

import (
	"sync"
)

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

var poolLike = sync.Pool{
	New: func() any {
		return new(Like)
	},
}

// GetLike() 从对象池中获取Like
func GetLike() *Like {
	return poolLike.Get().(*Like)
}

// ReleaseLike 释放Like
func ReleaseLike(v *Like) {
	v.Source = ""
	v.Id = ""
	v.Album = ""
	v.Author = ""
	v.Name = ""
	poolLike.Put(v)
}
