package alihouse

import (
	"sync"
)

// BrokerPointResultDto 结构体
type BrokerPointResultDto struct {
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBrokerPointResultDto = sync.Pool{
	New: func() any {
		return new(BrokerPointResultDto)
	},
}

// GetBrokerPointResultDto() 从对象池中获取BrokerPointResultDto
func GetBrokerPointResultDto() *BrokerPointResultDto {
	return poolBrokerPointResultDto.Get().(*BrokerPointResultDto)
}

// ReleaseBrokerPointResultDto 释放BrokerPointResultDto
func ReleaseBrokerPointResultDto(v *BrokerPointResultDto) {
	v.OuterId = ""
	v.Msg = ""
	v.Success = false
	poolBrokerPointResultDto.Put(v)
}
