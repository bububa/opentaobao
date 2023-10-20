package tvupadmin

import (
	"sync"
)

// YunosOsupdateOsfotaPublishResult 结构体
type YunosOsupdateOsfotaPublishResult struct {
	// 是否操作成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolYunosOsupdateOsfotaPublishResult = sync.Pool{
	New: func() any {
		return new(YunosOsupdateOsfotaPublishResult)
	},
}

// GetYunosOsupdateOsfotaPublishResult() 从对象池中获取YunosOsupdateOsfotaPublishResult
func GetYunosOsupdateOsfotaPublishResult() *YunosOsupdateOsfotaPublishResult {
	return poolYunosOsupdateOsfotaPublishResult.Get().(*YunosOsupdateOsfotaPublishResult)
}

// ReleaseYunosOsupdateOsfotaPublishResult 释放YunosOsupdateOsfotaPublishResult
func ReleaseYunosOsupdateOsfotaPublishResult(v *YunosOsupdateOsfotaPublishResult) {
	v.Data = false
	poolYunosOsupdateOsfotaPublishResult.Put(v)
}
