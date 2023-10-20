package tmallgeniescp

import (
	"sync"
)

// ConsensusDemandParamDto 结构体
type ConsensusDemandParamDto struct {
	// 物料号
	PrdId string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
	// 渠道
	CustId string `json:"cust_id,omitempty" xml:"cust_id,omitempty"`
	// 共识需求数量
	ConsensusDemand string `json:"consensus_demand,omitempty" xml:"consensus_demand,omitempty"`
	// 关键值日期
	KeyFigureDate string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 元素扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
}

var poolConsensusDemandParamDto = sync.Pool{
	New: func() any {
		return new(ConsensusDemandParamDto)
	},
}

// GetConsensusDemandParamDto() 从对象池中获取ConsensusDemandParamDto
func GetConsensusDemandParamDto() *ConsensusDemandParamDto {
	return poolConsensusDemandParamDto.Get().(*ConsensusDemandParamDto)
}

// ReleaseConsensusDemandParamDto 释放ConsensusDemandParamDto
func ReleaseConsensusDemandParamDto(v *ConsensusDemandParamDto) {
	v.PrdId = ""
	v.CustId = ""
	v.ConsensusDemand = ""
	v.KeyFigureDate = ""
	v.Tenant = ""
	v.ExtendJson = ""
	poolConsensusDemandParamDto.Put(v)
}
