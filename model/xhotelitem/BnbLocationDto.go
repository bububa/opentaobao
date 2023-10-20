package xhotelitem

import (
	"sync"
)

// BnbLocationDto 结构体
type BnbLocationDto struct {
	// domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 门店地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 商圈
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 坐标类型，现在支持：G – Google; B – 百度; A – 高德; M – Mapbar; L – 灵图
	PositionType string `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 门店英文地址
	EnAddress string `json:"en_address,omitempty" xml:"en_address,omitempty"`
	// 接待地址
	ReceptionAddress string `json:"reception_address,omitempty" xml:"reception_address,omitempty"`
	// 时区0到+11或者0到-11
	Timezone string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// 门牌号
	Doorplate string `json:"doorplate,omitempty" xml:"doorplate,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 城市名称，优先取city字段，city字段如果为空会校验cityName
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新时为可选）
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 国别 0:国内;1:国外。默认是国内
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 省份编码http://hotel.alitrip.com/area.htm
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
}

var poolBnbLocationDto = sync.Pool{
	New: func() any {
		return new(BnbLocationDto)
	},
}

// GetBnbLocationDto() 从对象池中获取BnbLocationDto
func GetBnbLocationDto() *BnbLocationDto {
	return poolBnbLocationDto.Get().(*BnbLocationDto)
}

// ReleaseBnbLocationDto 释放BnbLocationDto
func ReleaseBnbLocationDto(v *BnbLocationDto) {
	v.Country = ""
	v.Address = ""
	v.Business = ""
	v.Latitude = ""
	v.PositionType = ""
	v.EnAddress = ""
	v.ReceptionAddress = ""
	v.Timezone = ""
	v.Doorplate = ""
	v.Longitude = ""
	v.CityName = ""
	v.City = 0
	v.Domestic = 0
	v.Province = 0
	v.District = 0
	poolBnbLocationDto.Put(v)
}
