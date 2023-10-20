package tvupadmin

import (
	"sync"
)

// YunosOsupdateAppversionPublishResult 结构体
type YunosOsupdateAppversionPublishResult struct {
	// 操作是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolYunosOsupdateAppversionPublishResult = sync.Pool{
	New: func() any {
		return new(YunosOsupdateAppversionPublishResult)
	},
}

// GetYunosOsupdateAppversionPublishResult() 从对象池中获取YunosOsupdateAppversionPublishResult
func GetYunosOsupdateAppversionPublishResult() *YunosOsupdateAppversionPublishResult {
	return poolYunosOsupdateAppversionPublishResult.Get().(*YunosOsupdateAppversionPublishResult)
}

// ReleaseYunosOsupdateAppversionPublishResult 释放YunosOsupdateAppversionPublishResult
func ReleaseYunosOsupdateAppversionPublishResult(v *YunosOsupdateAppversionPublishResult) {
	v.Data = false
	poolYunosOsupdateAppversionPublishResult.Put(v)
}
