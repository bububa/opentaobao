package axindata

import (
	"sync"
)

// FscTripDivisionApiDto 结构体
type FscTripDivisionApiDto struct {
	// 行政区划名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 行政区划英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 目的地ID
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// 行政区划级别
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 行政区划父级ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 是否境外
	Abroad bool `json:"abroad,omitempty" xml:"abroad,omitempty"`
}

var poolFscTripDivisionApiDto = sync.Pool{
	New: func() any {
		return new(FscTripDivisionApiDto)
	},
}

// GetFscTripDivisionApiDto() 从对象池中获取FscTripDivisionApiDto
func GetFscTripDivisionApiDto() *FscTripDivisionApiDto {
	return poolFscTripDivisionApiDto.Get().(*FscTripDivisionApiDto)
}

// ReleaseFscTripDivisionApiDto 释放FscTripDivisionApiDto
func ReleaseFscTripDivisionApiDto(v *FscTripDivisionApiDto) {
	v.Name = ""
	v.NameEn = ""
	v.CountryName = ""
	v.DivisionId = 0
	v.Level = 0
	v.ParentId = 0
	v.Abroad = false
	poolFscTripDivisionApiDto.Put(v)
}
