package tbk

import (
	"sync"
)

// VegasCpaReportDto 结构体
type VegasCpaReportDto struct {
	// 奖励金额；按入参是预估/结算，区分获得金额为预估or可结算结果；
	RewardAmount string `json:"reward_amount,omitempty" xml:"reward_amount,omitempty"`
	// 统计日期（统计活动期内，截止 统计日期 的数据）
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 媒体三段式id，当查询数据为pid维度时返回该字段
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 活动相关数据信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 符合奖励要求的累计用户数；按入参是预估/结算，区分用户数为预估or可结算结果；
	Union30dLxUv int64 `json:"union_30d_lx_uv,omitempty" xml:"union_30d_lx_uv,omitempty"`
	// rid，当查询数据为rid维度时返回该字段
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 数据类型:1预估 2结算
	QueryType int64 `json:"query_type,omitempty" xml:"query_type,omitempty"`
}

var poolVegasCpaReportDto = sync.Pool{
	New: func() any {
		return new(VegasCpaReportDto)
	},
}

// GetVegasCpaReportDto() 从对象池中获取VegasCpaReportDto
func GetVegasCpaReportDto() *VegasCpaReportDto {
	return poolVegasCpaReportDto.Get().(*VegasCpaReportDto)
}

// ReleaseVegasCpaReportDto 释放VegasCpaReportDto
func ReleaseVegasCpaReportDto(v *VegasCpaReportDto) {
	v.RewardAmount = ""
	v.BizDate = ""
	v.Pid = ""
	v.ExtInfo = ""
	v.Union30dLxUv = 0
	v.RelationId = 0
	v.QueryType = 0
	poolVegasCpaReportDto.Put(v)
}
