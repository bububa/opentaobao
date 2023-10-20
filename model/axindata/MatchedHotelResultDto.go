package axindata

import (
	"sync"
)

// MatchedHotelResultDto 结构体
type MatchedHotelResultDto struct {
	// 酒店匹配结果列表，会有相应分值
	MatchedHotelDataList []MatchedHotelDataDto `json:"matched_hotel_data_list,omitempty" xml:"matched_hotel_data_list>matched_hotel_data_dto,omitempty"`
	// 参数酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 错误code
	ErCode string `json:"er_code,omitempty" xml:"er_code,omitempty"`
	// 错误信息
	ErMsg string `json:"er_msg,omitempty" xml:"er_msg,omitempty"`
	// 当前匹配是否有异常
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolMatchedHotelResultDto = sync.Pool{
	New: func() any {
		return new(MatchedHotelResultDto)
	},
}

// GetMatchedHotelResultDto() 从对象池中获取MatchedHotelResultDto
func GetMatchedHotelResultDto() *MatchedHotelResultDto {
	return poolMatchedHotelResultDto.Get().(*MatchedHotelResultDto)
}

// ReleaseMatchedHotelResultDto 释放MatchedHotelResultDto
func ReleaseMatchedHotelResultDto(v *MatchedHotelResultDto) {
	v.MatchedHotelDataList = v.MatchedHotelDataList[:0]
	v.HotelName = ""
	v.ErCode = ""
	v.ErMsg = ""
	v.Succ = false
	poolMatchedHotelResultDto.Put(v)
}
