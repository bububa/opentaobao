package alitrippoi

import (
	"sync"
)

// AlitripPlatformPoiRawPoioutResult 结构体
type AlitripPlatformPoiRawPoioutResult struct {
	// 返回素材id
	Datas []AlitripPlatformPoiRawPoioutData `json:"datas,omitempty" xml:"datas>alitrip_platform_poi_raw_poiout_data,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数(不可用)
	TotalRecords int64 `json:"total_records,omitempty" xml:"total_records,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPlatformPoiRawPoioutResult = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawPoioutResult)
	},
}

// GetAlitripPlatformPoiRawPoioutResult() 从对象池中获取AlitripPlatformPoiRawPoioutResult
func GetAlitripPlatformPoiRawPoioutResult() *AlitripPlatformPoiRawPoioutResult {
	return poolAlitripPlatformPoiRawPoioutResult.Get().(*AlitripPlatformPoiRawPoioutResult)
}

// ReleaseAlitripPlatformPoiRawPoioutResult 释放AlitripPlatformPoiRawPoioutResult
func ReleaseAlitripPlatformPoiRawPoioutResult(v *AlitripPlatformPoiRawPoioutResult) {
	v.Datas = v.Datas[:0]
	v.ResultCode = ""
	v.Message = ""
	v.TotalRecords = 0
	v.Success = false
	poolAlitripPlatformPoiRawPoioutResult.Put(v)
}
