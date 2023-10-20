package icburfq

import (
	"sync"
)

// QueryDto 结构体
type QueryDto struct {
	// 系统参数qn-homepage
	Site string `json:"site,omitempty" xml:"site,omitempty"`
	// 系统参数U_P_I
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 推荐数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 当前页面数
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolQueryDto = sync.Pool{
	New: func() any {
		return new(QueryDto)
	},
}

// GetQueryDto() 从对象池中获取QueryDto
func GetQueryDto() *QueryDto {
	return poolQueryDto.Get().(*QueryDto)
}

// ReleaseQueryDto 释放QueryDto
func ReleaseQueryDto(v *QueryDto) {
	v.Site = ""
	v.Type = ""
	v.Count = 0
	v.Current = 0
	v.PageSize = 0
	poolQueryDto.Put(v)
}
