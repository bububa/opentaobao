package icbu

import (
	"sync"
)

// PhotoGroupResult 结构体
type PhotoGroupResult struct {
	// add操作中表示新增的图片分组，rename操作中表示重命名的分组，delete操作中返回分组信息
	PhotobankGroup *PhotobankGroup `json:"photobank_group,omitempty" xml:"photobank_group,omitempty"`
}

var poolPhotoGroupResult = sync.Pool{
	New: func() any {
		return new(PhotoGroupResult)
	},
}

// GetPhotoGroupResult() 从对象池中获取PhotoGroupResult
func GetPhotoGroupResult() *PhotoGroupResult {
	return poolPhotoGroupResult.Get().(*PhotoGroupResult)
}

// ReleasePhotoGroupResult 释放PhotoGroupResult
func ReleasePhotoGroupResult(v *PhotoGroupResult) {
	v.PhotobankGroup = nil
	poolPhotoGroupResult.Put(v)
}
