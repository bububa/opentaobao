package tvupadmin

import (
	"sync"
)

// DeviceAdapterDto 结构体
type DeviceAdapterDto struct {
	// brandName
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// highestSystemVersion
	HighestSystemVersion string `json:"highest_system_version,omitempty" xml:"highest_system_version,omitempty"`
	// minimumSystemVersion
	MinimumSystemVersion string `json:"minimum_system_version,omitempty" xml:"minimum_system_version,omitempty"`
	// modelName
	ModelName string `json:"model_name,omitempty" xml:"model_name,omitempty"`
	// realTypeName
	RealTypeName string `json:"real_type_name,omitempty" xml:"real_type_name,omitempty"`
	// brandId
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// modelId
	ModelId int64 `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// realTypeId
	RealTypeId int64 `json:"real_type_id,omitempty" xml:"real_type_id,omitempty"`
}

var poolDeviceAdapterDto = sync.Pool{
	New: func() any {
		return new(DeviceAdapterDto)
	},
}

// GetDeviceAdapterDto() 从对象池中获取DeviceAdapterDto
func GetDeviceAdapterDto() *DeviceAdapterDto {
	return poolDeviceAdapterDto.Get().(*DeviceAdapterDto)
}

// ReleaseDeviceAdapterDto 释放DeviceAdapterDto
func ReleaseDeviceAdapterDto(v *DeviceAdapterDto) {
	v.BrandName = ""
	v.HighestSystemVersion = ""
	v.MinimumSystemVersion = ""
	v.ModelName = ""
	v.RealTypeName = ""
	v.BrandId = 0
	v.ModelId = 0
	v.RealTypeId = 0
	poolDeviceAdapterDto.Put(v)
}
