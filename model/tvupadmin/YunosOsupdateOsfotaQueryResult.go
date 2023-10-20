package tvupadmin

import (
	"sync"
)

// YunosOsupdateOsfotaQueryResult 结构体
type YunosOsupdateOsfotaQueryResult struct {
	// 具体内容
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

var poolYunosOsupdateOsfotaQueryResult = sync.Pool{
	New: func() any {
		return new(YunosOsupdateOsfotaQueryResult)
	},
}

// GetYunosOsupdateOsfotaQueryResult() 从对象池中获取YunosOsupdateOsfotaQueryResult
func GetYunosOsupdateOsfotaQueryResult() *YunosOsupdateOsfotaQueryResult {
	return poolYunosOsupdateOsfotaQueryResult.Get().(*YunosOsupdateOsfotaQueryResult)
}

// ReleaseYunosOsupdateOsfotaQueryResult 释放YunosOsupdateOsfotaQueryResult
func ReleaseYunosOsupdateOsfotaQueryResult(v *YunosOsupdateOsfotaQueryResult) {
	v.Data = ""
	poolYunosOsupdateOsfotaQueryResult.Put(v)
}
