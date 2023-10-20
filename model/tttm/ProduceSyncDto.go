package tttm

import (
	"sync"
)

// ProduceSyncDto 结构体
type ProduceSyncDto struct {
	// 生产节点
	ProduceLink string `json:"produce_link,omitempty" xml:"produce_link,omitempty"`
	// 提交时间
	SubmitTime string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	// 实际时间
	FinishTime string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	// 顺序
	LinkSort int64 `json:"link_sort,omitempty" xml:"link_sort,omitempty"`
	// 产量
	ProduceNum int64 `json:"produce_num,omitempty" xml:"produce_num,omitempty"`
	// 次品量
	DefectiveNum int64 `json:"defective_num,omitempty" xml:"defective_num,omitempty"`
	// 生产状态
	ProduceStatus int64 `json:"produce_status,omitempty" xml:"produce_status,omitempty"`
}

var poolProduceSyncDto = sync.Pool{
	New: func() any {
		return new(ProduceSyncDto)
	},
}

// GetProduceSyncDto() 从对象池中获取ProduceSyncDto
func GetProduceSyncDto() *ProduceSyncDto {
	return poolProduceSyncDto.Get().(*ProduceSyncDto)
}

// ReleaseProduceSyncDto 释放ProduceSyncDto
func ReleaseProduceSyncDto(v *ProduceSyncDto) {
	v.ProduceLink = ""
	v.SubmitTime = ""
	v.FinishTime = ""
	v.LinkSort = 0
	v.ProduceNum = 0
	v.DefectiveNum = 0
	v.ProduceStatus = 0
	poolProduceSyncDto.Put(v)
}
