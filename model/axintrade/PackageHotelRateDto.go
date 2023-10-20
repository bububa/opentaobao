package axintrade

import (
	"sync"
)

// PackageHotelRateDto 结构体
type PackageHotelRateDto struct {
	// 酒店床型名称
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 酒店rate信息
	HotelRate *HotelRateDto `json:"hotel_rate,omitempty" xml:"hotel_rate,omitempty"`
	// 酒店间夜数
	RoomNightNum int64 `json:"room_night_num,omitempty" xml:"room_night_num,omitempty"`
	// 成人数
	AdultNum int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
	// 儿童数
	ChildrenNum int64 `json:"children_num,omitempty" xml:"children_num,omitempty"`
}

var poolPackageHotelRateDto = sync.Pool{
	New: func() any {
		return new(PackageHotelRateDto)
	},
}

// GetPackageHotelRateDto() 从对象池中获取PackageHotelRateDto
func GetPackageHotelRateDto() *PackageHotelRateDto {
	return poolPackageHotelRateDto.Get().(*PackageHotelRateDto)
}

// ReleasePackageHotelRateDto 释放PackageHotelRateDto
func ReleasePackageHotelRateDto(v *PackageHotelRateDto) {
	v.BedType = ""
	v.HotelRate = nil
	v.RoomNightNum = 0
	v.AdultNum = 0
	v.ChildrenNum = 0
	poolPackageHotelRateDto.Put(v)
}
