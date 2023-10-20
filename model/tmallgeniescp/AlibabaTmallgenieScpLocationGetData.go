package tmallgeniescp

import (
	"sync"
)

// AlibabaTmallgenieScpLocationGetData 结构体
type AlibabaTmallgenieScpLocationGetData struct {
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// CDC或者RDC类型
	LocationType string `json:"location_type,omitempty" xml:"location_type,omitempty"`
	// CDC或者RDC名称
	LocationName string `json:"location_name,omitempty" xml:"location_name,omitempty"`
	// 地点编码
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
}

var poolAlibabaTmallgenieScpLocationGetData = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpLocationGetData)
	},
}

// GetAlibabaTmallgenieScpLocationGetData() 从对象池中获取AlibabaTmallgenieScpLocationGetData
func GetAlibabaTmallgenieScpLocationGetData() *AlibabaTmallgenieScpLocationGetData {
	return poolAlibabaTmallgenieScpLocationGetData.Get().(*AlibabaTmallgenieScpLocationGetData)
}

// ReleaseAlibabaTmallgenieScpLocationGetData 释放AlibabaTmallgenieScpLocationGetData
func ReleaseAlibabaTmallgenieScpLocationGetData(v *AlibabaTmallgenieScpLocationGetData) {
	v.ExtendJson = ""
	v.Tenant = ""
	v.LocationType = ""
	v.LocationName = ""
	v.LocationCode = ""
	poolAlibabaTmallgenieScpLocationGetData.Put(v)
}
