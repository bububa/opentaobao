package traderate

import (
	"sync"
)

// TmallRateTagDetail 结构体
type TmallRateTagDetail struct {
	// 标签名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 反应该标签的评价数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 标签的极性：1正极 -1负极
	Posi bool `json:"posi,omitempty" xml:"posi,omitempty"`
}

var poolTmallRateTagDetail = sync.Pool{
	New: func() any {
		return new(TmallRateTagDetail)
	},
}

// GetTmallRateTagDetail() 从对象池中获取TmallRateTagDetail
func GetTmallRateTagDetail() *TmallRateTagDetail {
	return poolTmallRateTagDetail.Get().(*TmallRateTagDetail)
}

// ReleaseTmallRateTagDetail 释放TmallRateTagDetail
func ReleaseTmallRateTagDetail(v *TmallRateTagDetail) {
	v.TagName = ""
	v.Count = 0
	v.Posi = false
	poolTmallRateTagDetail.Put(v)
}
