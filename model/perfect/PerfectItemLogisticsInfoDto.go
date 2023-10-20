package perfect

import (
	"sync"
)

// PerfectItemLogisticsInfoDto 结构体
type PerfectItemLogisticsInfoDto struct {
	// 城市编码,默认北京
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 运费模板ID
	PostageTemplateCode string `json:"postage_template_code,omitempty" xml:"postage_template_code,omitempty"`
	// 省份编码,默认北京
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 门店库ID
	StoreGroupCode string `json:"store_group_code,omitempty" xml:"store_group_code,omitempty"`
}

var poolPerfectItemLogisticsInfoDto = sync.Pool{
	New: func() any {
		return new(PerfectItemLogisticsInfoDto)
	},
}

// GetPerfectItemLogisticsInfoDto() 从对象池中获取PerfectItemLogisticsInfoDto
func GetPerfectItemLogisticsInfoDto() *PerfectItemLogisticsInfoDto {
	return poolPerfectItemLogisticsInfoDto.Get().(*PerfectItemLogisticsInfoDto)
}

// ReleasePerfectItemLogisticsInfoDto 释放PerfectItemLogisticsInfoDto
func ReleasePerfectItemLogisticsInfoDto(v *PerfectItemLogisticsInfoDto) {
	v.CityCode = ""
	v.PostageTemplateCode = ""
	v.ProvinceCode = ""
	v.StoreGroupCode = ""
	poolPerfectItemLogisticsInfoDto.Put(v)
}
