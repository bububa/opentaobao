package tblogistics

import (
	"sync"
)

// TopLogisticsNodeDto 结构体
type TopLogisticsNodeDto struct {
	// 节点描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 节点枚举
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 当前节点发生时间
	StatusTime int64 `json:"status_time,omitempty" xml:"status_time,omitempty"`
}

var poolTopLogisticsNodeDto = sync.Pool{
	New: func() any {
		return new(TopLogisticsNodeDto)
	},
}

// GetTopLogisticsNodeDto() 从对象池中获取TopLogisticsNodeDto
func GetTopLogisticsNodeDto() *TopLogisticsNodeDto {
	return poolTopLogisticsNodeDto.Get().(*TopLogisticsNodeDto)
}

// ReleaseTopLogisticsNodeDto 释放TopLogisticsNodeDto
func ReleaseTopLogisticsNodeDto(v *TopLogisticsNodeDto) {
	v.StatusDesc = ""
	v.Action = ""
	v.StatusTime = 0
	poolTopLogisticsNodeDto.Put(v)
}
