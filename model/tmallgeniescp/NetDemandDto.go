package tmallgeniescp

import (
	"sync"
)

// NetDemandDto 结构体
type NetDemandDto struct {
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 关键日期值
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// 地点
	Locid string `json:"locid,omitempty" xml:"locid,omitempty"`
	// 物料号
	Prdid string `json:"prdid,omitempty" xml:"prdid,omitempty"`
	// 净需求值
	NetDemand int64 `json:"net_demand,omitempty" xml:"net_demand,omitempty"`
}

var poolNetDemandDto = sync.Pool{
	New: func() any {
		return new(NetDemandDto)
	},
}

// GetNetDemandDto() 从对象池中获取NetDemandDto
func GetNetDemandDto() *NetDemandDto {
	return poolNetDemandDto.Get().(*NetDemandDto)
}

// ReleaseNetDemandDto 释放NetDemandDto
func ReleaseNetDemandDto(v *NetDemandDto) {
	v.ExtendJson = ""
	v.Tenant = ""
	v.KeyFigureDate = ""
	v.Locid = ""
	v.Prdid = ""
	v.NetDemand = 0
	poolNetDemandDto.Put(v)
}
