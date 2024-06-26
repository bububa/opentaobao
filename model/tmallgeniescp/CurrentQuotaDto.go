package tmallgeniescp

import (
	"sync"
)

// CurrentQuotaDto 结构体
type CurrentQuotaDto struct {
	// 关键日期值
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// 比例（精确到2位小数：0.34）
	Ratio string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// 二级物料-物料编码
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 二级物料供应商-locCode
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

var poolCurrentQuotaDto = sync.Pool{
	New: func() any {
		return new(CurrentQuotaDto)
	},
}

// GetCurrentQuotaDto() 从对象池中获取CurrentQuotaDto
func GetCurrentQuotaDto() *CurrentQuotaDto {
	return poolCurrentQuotaDto.Get().(*CurrentQuotaDto)
}

// ReleaseCurrentQuotaDto 释放CurrentQuotaDto
func ReleaseCurrentQuotaDto(v *CurrentQuotaDto) {
	v.KeyFigureDate = ""
	v.Ratio = ""
	v.MaterielCode = ""
	v.LocationCode = ""
	v.ExtendJson = ""
	v.Tenant = ""
	poolCurrentQuotaDto.Put(v)
}
