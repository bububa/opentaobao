package kbalgo

import (
	"sync"
)

// AlscPoiToBaiduResult 结构体
type AlscPoiToBaiduResult struct {
	// datas
	Datas []PoiToBaiduData `json:"datas,omitempty" xml:"datas>poi_to_baidu_data,omitempty"`
	// 附加信息或错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// poi总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 状态码：0-success，1-fail
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAlscPoiToBaiduResult = sync.Pool{
	New: func() any {
		return new(AlscPoiToBaiduResult)
	},
}

// GetAlscPoiToBaiduResult() 从对象池中获取AlscPoiToBaiduResult
func GetAlscPoiToBaiduResult() *AlscPoiToBaiduResult {
	return poolAlscPoiToBaiduResult.Get().(*AlscPoiToBaiduResult)
}

// ReleaseAlscPoiToBaiduResult 释放AlscPoiToBaiduResult
func ReleaseAlscPoiToBaiduResult(v *AlscPoiToBaiduResult) {
	v.Datas = v.Datas[:0]
	v.Message = ""
	v.Total = 0
	v.Status = 0
	poolAlscPoiToBaiduResult.Put(v)
}
