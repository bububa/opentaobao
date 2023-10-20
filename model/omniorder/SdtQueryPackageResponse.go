package omniorder

import (
	"sync"
)

// SdtQueryPackageResponse 结构体
type SdtQueryPackageResponse struct {
	// 站点信息
	Stations []SdtStationDto `json:"stations,omitempty" xml:"stations>sdt_station_dto,omitempty"`
}

var poolSdtQueryPackageResponse = sync.Pool{
	New: func() any {
		return new(SdtQueryPackageResponse)
	},
}

// GetSdtQueryPackageResponse() 从对象池中获取SdtQueryPackageResponse
func GetSdtQueryPackageResponse() *SdtQueryPackageResponse {
	return poolSdtQueryPackageResponse.Get().(*SdtQueryPackageResponse)
}

// ReleaseSdtQueryPackageResponse 释放SdtQueryPackageResponse
func ReleaseSdtQueryPackageResponse(v *SdtQueryPackageResponse) {
	v.Stations = v.Stations[:0]
	poolSdtQueryPackageResponse.Put(v)
}
