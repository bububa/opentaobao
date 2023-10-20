package util

import (
	"sync"
)

// CountryDto 结构体
type CountryDto struct {
	// 国家ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 国家名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 国家CODE
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolCountryDto = sync.Pool{
	New: func() any {
		return new(CountryDto)
	},
}

// GetCountryDto() 从对象池中获取CountryDto
func GetCountryDto() *CountryDto {
	return poolCountryDto.Get().(*CountryDto)
}

// ReleaseCountryDto 释放CountryDto
func ReleaseCountryDto(v *CountryDto) {
	v.Id = ""
	v.Name = ""
	v.Code = ""
	poolCountryDto.Put(v)
}
