package tvupadmin

import (
	"sync"
)

// YunosOsupdateOsfotaAddResult 结构体
type YunosOsupdateOsfotaAddResult struct {
	// 是否操作成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolYunosOsupdateOsfotaAddResult = sync.Pool{
	New: func() any {
		return new(YunosOsupdateOsfotaAddResult)
	},
}

// GetYunosOsupdateOsfotaAddResult() 从对象池中获取YunosOsupdateOsfotaAddResult
func GetYunosOsupdateOsfotaAddResult() *YunosOsupdateOsfotaAddResult {
	return poolYunosOsupdateOsfotaAddResult.Get().(*YunosOsupdateOsfotaAddResult)
}

// ReleaseYunosOsupdateOsfotaAddResult 释放YunosOsupdateOsfotaAddResult
func ReleaseYunosOsupdateOsfotaAddResult(v *YunosOsupdateOsfotaAddResult) {
	v.Data = false
	poolYunosOsupdateOsfotaAddResult.Put(v)
}
