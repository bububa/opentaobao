package tblogistics

import (
	"sync"
)

// TopLogisticsMailDto 结构体
type TopLogisticsMailDto struct {
	// 物流节点列表
	TraceList []TopLogisticsNodeDto `json:"trace_list,omitempty" xml:"trace_list>top_logistics_node_dto,omitempty"`
	// 运单号
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 物流公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 当前最新节点
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}

var poolTopLogisticsMailDto = sync.Pool{
	New: func() any {
		return new(TopLogisticsMailDto)
	},
}

// GetTopLogisticsMailDto() 从对象池中获取TopLogisticsMailDto
func GetTopLogisticsMailDto() *TopLogisticsMailDto {
	return poolTopLogisticsMailDto.Get().(*TopLogisticsMailDto)
}

// ReleaseTopLogisticsMailDto 释放TopLogisticsMailDto
func ReleaseTopLogisticsMailDto(v *TopLogisticsMailDto) {
	v.TraceList = v.TraceList[:0]
	v.OutSid = ""
	v.CompanyName = ""
	v.Status = ""
	v.Tid = 0
	poolTopLogisticsMailDto.Put(v)
}
