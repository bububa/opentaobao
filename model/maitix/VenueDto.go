package maitix

import (
	"sync"
)

// VenueDto 结构体
type VenueDto struct {
	// 场馆名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 场馆地址
	VenueAddress string `json:"venue_address,omitempty" xml:"venue_address,omitempty"`
	// 场馆id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolVenueDto = sync.Pool{
	New: func() any {
		return new(VenueDto)
	},
}

// GetVenueDto() 从对象池中获取VenueDto
func GetVenueDto() *VenueDto {
	return poolVenueDto.Get().(*VenueDto)
}

// ReleaseVenueDto 释放VenueDto
func ReleaseVenueDto(v *VenueDto) {
	v.Name = ""
	v.Lng = ""
	v.Lat = ""
	v.VenueAddress = ""
	v.Id = 0
	poolVenueDto.Put(v)
}
