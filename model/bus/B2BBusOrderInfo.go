package bus

import (
	"sync"
)

// B2BBusOrderInfo 结构体
type B2BBusOrderInfo struct {
	// 票信息
	BusB2bTicketInfoList []B2BTicketInfo `json:"bus_b2b_ticket_info_list,omitempty" xml:"bus_b2b_ticket_info_list>b2b_ticket_info,omitempty"`
	// 支付宝交易流水号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 出票成功时间
	IssueTime string `json:"issue_time,omitempty" xml:"issue_time,omitempty"`
	// 车次对象
	B2BBusLineInfo *B2BBusLineInfo `json:"b2_b_bus_line_info,omitempty" xml:"b2_b_bus_line_info,omitempty"`
	// 取票人信息
	B2BFetchHolderInfo *B2BFetchHolderInfo `json:"b2_b_fetch_holder_info,omitempty" xml:"b2_b_fetch_holder_info,omitempty"`
	// 取票信息
	B2BFetchTicket *B2BFetchTicket `json:"b2_b_fetch_ticket,omitempty" xml:"b2_b_fetch_ticket,omitempty"`
	// 阿里订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 总价
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

var poolB2BBusOrderInfo = sync.Pool{
	New: func() any {
		return new(B2BBusOrderInfo)
	},
}

// GetB2BBusOrderInfo() 从对象池中获取B2BBusOrderInfo
func GetB2BBusOrderInfo() *B2BBusOrderInfo {
	return poolB2BBusOrderInfo.Get().(*B2BBusOrderInfo)
}

// ReleaseB2BBusOrderInfo 释放B2BBusOrderInfo
func ReleaseB2BBusOrderInfo(v *B2BBusOrderInfo) {
	v.BusB2bTicketInfoList = v.BusB2bTicketInfoList[:0]
	v.AlipayTradeId = ""
	v.GmtCreate = ""
	v.IssueTime = ""
	v.B2BBusLineInfo = nil
	v.B2BFetchHolderInfo = nil
	v.B2BFetchTicket = nil
	v.MainOrderId = 0
	v.OrderStatus = 0
	v.TicketCount = 0
	v.TotalPrice = 0
	poolB2BBusOrderInfo.Put(v)
}
