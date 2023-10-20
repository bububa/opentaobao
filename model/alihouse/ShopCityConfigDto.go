package alihouse

import (
	"sync"
)

// ShopCityConfigDto 结构体
type ShopCityConfigDto struct {
	// 城市配置首字母
	CityInitial string `json:"city_initial,omitempty" xml:"city_initial,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 城市配置外部唯一ID
	OuterCityConfigId string `json:"outer_city_config_id,omitempty" xml:"outer_city_config_id,omitempty"`
	// 是否为默认城市
	IsDefaultCity int64 `json:"is_default_city,omitempty" xml:"is_default_city,omitempty"`
	// 城市ID
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 是否展示
	IsShow int64 `json:"is_show,omitempty" xml:"is_show,omitempty"`
}

var poolShopCityConfigDto = sync.Pool{
	New: func() any {
		return new(ShopCityConfigDto)
	},
}

// GetShopCityConfigDto() 从对象池中获取ShopCityConfigDto
func GetShopCityConfigDto() *ShopCityConfigDto {
	return poolShopCityConfigDto.Get().(*ShopCityConfigDto)
}

// ReleaseShopCityConfigDto 释放ShopCityConfigDto
func ReleaseShopCityConfigDto(v *ShopCityConfigDto) {
	v.CityInitial = ""
	v.CityName = ""
	v.OuterCityConfigId = ""
	v.IsDefaultCity = 0
	v.CityId = 0
	v.IsShow = 0
	poolShopCityConfigDto.Put(v)
}
