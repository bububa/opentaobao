package media

import (
	"sync"
)

// VideoSearchCondition2 结构体
type VideoSearchCondition2 struct {
	// 当前页数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolVideoSearchCondition2 = sync.Pool{
	New: func() any {
		return new(VideoSearchCondition2)
	},
}

// GetVideoSearchCondition2() 从对象池中获取VideoSearchCondition2
func GetVideoSearchCondition2() *VideoSearchCondition2 {
	return poolVideoSearchCondition2.Get().(*VideoSearchCondition2)
}

// ReleaseVideoSearchCondition2 释放VideoSearchCondition2
func ReleaseVideoSearchCondition2(v *VideoSearchCondition2) {
	v.CurrentPage = 0
	v.PageSize = 0
	poolVideoSearchCondition2.Put(v)
}
