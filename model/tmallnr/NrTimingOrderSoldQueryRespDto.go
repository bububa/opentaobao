package tmallnr

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
