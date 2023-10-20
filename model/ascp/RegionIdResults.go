package ascp

import (
	"sync"
)

// RegionIdResults 结构体
type RegionIdResults struct {
	// 经度（高德）
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 纬度（高德）
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 错误原因
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

var poolRegionIdResults = sync.Pool{
	New: func() any {
		return new(RegionIdResults)
	},
}

// GetRegionIdResults() 从对象池中获取RegionIdResults
func GetRegionIdResults() *RegionIdResults {
	return poolRegionIdResults.Get().(*RegionIdResults)
}

// ReleaseRegionIdResults 释放RegionIdResults
func ReleaseRegionIdResults(v *RegionIdResults) {
	v.Latitude = ""
	v.Longitude = ""
	v.ErrorMessage = ""
	poolRegionIdResults.Put(v)
}
