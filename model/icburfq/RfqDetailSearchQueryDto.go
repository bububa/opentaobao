package icburfq

import (
	"sync"
)

// RfqDetailSearchQueryDto 结构体
type RfqDetailSearchQueryDto struct {
	// RFQ ID
	RfqId string `json:"rfq_id,omitempty" xml:"rfq_id,omitempty"`
}

var poolRfqDetailSearchQueryDto = sync.Pool{
	New: func() any {
		return new(RfqDetailSearchQueryDto)
	},
}

// GetRfqDetailSearchQueryDto() 从对象池中获取RfqDetailSearchQueryDto
func GetRfqDetailSearchQueryDto() *RfqDetailSearchQueryDto {
	return poolRfqDetailSearchQueryDto.Get().(*RfqDetailSearchQueryDto)
}

// ReleaseRfqDetailSearchQueryDto 释放RfqDetailSearchQueryDto
func ReleaseRfqDetailSearchQueryDto(v *RfqDetailSearchQueryDto) {
	v.RfqId = ""
	poolRfqDetailSearchQueryDto.Put(v)
}
