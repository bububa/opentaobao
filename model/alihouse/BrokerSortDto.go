package alihouse

import (
	"sync"
)

// BrokerSortDto 结构体
type BrokerSortDto struct {
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 排序号
	SortNo int64 `json:"sort_no,omitempty" xml:"sort_no,omitempty"`
}

var poolBrokerSortDto = sync.Pool{
	New: func() any {
		return new(BrokerSortDto)
	},
}

// GetBrokerSortDto() 从对象池中获取BrokerSortDto
func GetBrokerSortDto() *BrokerSortDto {
	return poolBrokerSortDto.Get().(*BrokerSortDto)
}

// ReleaseBrokerSortDto 释放BrokerSortDto
func ReleaseBrokerSortDto(v *BrokerSortDto) {
	v.OuterConsultantId = ""
	v.SortNo = 0
	poolBrokerSortDto.Put(v)
}
