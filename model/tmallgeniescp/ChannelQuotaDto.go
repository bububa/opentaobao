package tmallgeniescp

import (
	"sync"
)

// ChannelQuotaDto 结构体
type ChannelQuotaDto struct {
	// 关键日期值
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// 比例（精确到2位小数：0.34）
	Ratio string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// 物料code
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
	// 地点code
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
	// 渠道ID
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
}

var poolChannelQuotaDto = sync.Pool{
	New: func() any {
		return new(ChannelQuotaDto)
	},
}

// GetChannelQuotaDto() 从对象池中获取ChannelQuotaDto
func GetChannelQuotaDto() *ChannelQuotaDto {
	return poolChannelQuotaDto.Get().(*ChannelQuotaDto)
}

// ReleaseChannelQuotaDto 释放ChannelQuotaDto
func ReleaseChannelQuotaDto(v *ChannelQuotaDto) {
	v.KeyFigureDate = ""
	v.Ratio = ""
	v.MaterielCode = ""
	v.LocationCode = ""
	v.ChannelId = ""
	v.ExtendJson = ""
	v.Tenant = ""
	poolChannelQuotaDto.Put(v)
}
