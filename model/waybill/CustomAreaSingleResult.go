package waybill

import (
	"sync"
)

// CustomAreaSingleResult 结构体
type CustomAreaSingleResult struct {
	// keys
	Keys []KeyResult `json:"keys,omitempty" xml:"keys>key_result,omitempty"`
	// 自定义区内容的URL
	CustomAreaUrl string `json:"custom_area_url,omitempty" xml:"custom_area_url,omitempty"`
	// 自定义区id
	CustomAreaId int64 `json:"custom_area_id,omitempty" xml:"custom_area_id,omitempty"`
}

var poolCustomAreaSingleResult = sync.Pool{
	New: func() any {
		return new(CustomAreaSingleResult)
	},
}

// GetCustomAreaSingleResult() 从对象池中获取CustomAreaSingleResult
func GetCustomAreaSingleResult() *CustomAreaSingleResult {
	return poolCustomAreaSingleResult.Get().(*CustomAreaSingleResult)
}

// ReleaseCustomAreaSingleResult 释放CustomAreaSingleResult
func ReleaseCustomAreaSingleResult(v *CustomAreaSingleResult) {
	v.Keys = v.Keys[:0]
	v.CustomAreaUrl = ""
	v.CustomAreaId = 0
	poolCustomAreaSingleResult.Put(v)
}
