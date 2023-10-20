package xhotelitem

import (
	"sync"
)

// ReviewPicInfo 结构体
type ReviewPicInfo struct {
	// 图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolReviewPicInfo = sync.Pool{
	New: func() any {
		return new(ReviewPicInfo)
	},
}

// GetReviewPicInfo() 从对象池中获取ReviewPicInfo
func GetReviewPicInfo() *ReviewPicInfo {
	return poolReviewPicInfo.Get().(*ReviewPicInfo)
}

// ReleaseReviewPicInfo 释放ReviewPicInfo
func ReleaseReviewPicInfo(v *ReviewPicInfo) {
	v.Url = ""
	poolReviewPicInfo.Put(v)
}
