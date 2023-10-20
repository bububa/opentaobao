package alihouse

import (
	"sync"
)

// TradeToolBindResultDto 结构体
type TradeToolBindResultDto struct {
	// 外部工具唯一id，如购房登记ID
	OuterToolId string `json:"outer_tool_id,omitempty" xml:"outer_tool_id,omitempty"`
	// 外部楼盘和货ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部货id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTradeToolBindResultDto = sync.Pool{
	New: func() any {
		return new(TradeToolBindResultDto)
	},
}

// GetTradeToolBindResultDto() 从对象池中获取TradeToolBindResultDto
func GetTradeToolBindResultDto() *TradeToolBindResultDto {
	return poolTradeToolBindResultDto.Get().(*TradeToolBindResultDto)
}

// ReleaseTradeToolBindResultDto 释放TradeToolBindResultDto
func ReleaseTradeToolBindResultDto(v *TradeToolBindResultDto) {
	v.OuterToolId = ""
	v.OuterId = ""
	v.OuterTid = ""
	v.Code = ""
	v.Message = ""
	poolTradeToolBindResultDto.Put(v)
}
