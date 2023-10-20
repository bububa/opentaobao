package tmallnr

import (
	"sync"
)

// NrTimingOrderSoldQueryRespDto 结构体
type NrTimingOrderSoldQueryRespDto struct {
	// 主订单列表
	TradeOrderDetailDTOs []TradeOrderDetailDto `json:"trade_order_detail_d_t_os,omitempty" xml:"trade_order_detail_d_t_os>trade_order_detail_dto,omitempty"`
	// pageNo
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// totalNum
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolNrTimingOrderSoldQueryRespDto = sync.Pool{
	New: func() any {
		return new(NrTimingOrderSoldQueryRespDto)
	},
}

// GetNrTimingOrderSoldQueryRespDto() 从对象池中获取NrTimingOrderSoldQueryRespDto
func GetNrTimingOrderSoldQueryRespDto() *NrTimingOrderSoldQueryRespDto {
	return poolNrTimingOrderSoldQueryRespDto.Get().(*NrTimingOrderSoldQueryRespDto)
}

// ReleaseNrTimingOrderSoldQueryRespDto 释放NrTimingOrderSoldQueryRespDto
func ReleaseNrTimingOrderSoldQueryRespDto(v *NrTimingOrderSoldQueryRespDto) {
	v.TradeOrderDetailDTOs = v.TradeOrderDetailDTOs[:0]
	v.PageNo = 0
	v.TotalNum = 0
	v.PageSize = 0
	poolNrTimingOrderSoldQueryRespDto.Put(v)
}
