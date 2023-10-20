package btrip

import (
	"sync"
)

// OpenApiTrainOrderDetailRSV2 结构体
type OpenApiTrainOrderDetailRSV2 struct {
	// 改签票列表
	ChangeTicketInfoList []TrainChangeTicketInfo `json:"change_ticket_info_list,omitempty" xml:"change_ticket_info_list>train_change_ticket_info,omitempty"`
	// 退票列表
	RefundTicketInfoList []TrainRefundTicketInfo `json:"refund_ticket_info_list,omitempty" xml:"refund_ticket_info_list>train_refund_ticket_info,omitempty"`
	// 出行人列表
	PassengerInfoList []PassengerInfo `json:"passenger_info_list,omitempty" xml:"passenger_info_list>passenger_info,omitempty"`
	// 订单费用列表
	PriceInfoList []PriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>price_info,omitempty"`
	// 订单基础信息
	OrderBaseInfo *OrderBaseInfo `json:"order_base_info,omitempty" xml:"order_base_info,omitempty"`
	// 发票信息
	InvoiceInfo *InvoiceInfo `json:"invoice_info,omitempty" xml:"invoice_info,omitempty"`
	// 订单信息
	TrainOrderInfo *TrainOrderInfo `json:"trainOrderInfo,omitempty" xml:"trainOrderInfo,omitempty"`
}

var poolOpenApiTrainOrderDetailRSV2 = sync.Pool{
	New: func() any {
		return new(OpenApiTrainOrderDetailRSV2)
	},
}

// GetOpenApiTrainOrderDetailRSV2() 从对象池中获取OpenApiTrainOrderDetailRSV2
func GetOpenApiTrainOrderDetailRSV2() *OpenApiTrainOrderDetailRSV2 {
	return poolOpenApiTrainOrderDetailRSV2.Get().(*OpenApiTrainOrderDetailRSV2)
}

// ReleaseOpenApiTrainOrderDetailRSV2 释放OpenApiTrainOrderDetailRSV2
func ReleaseOpenApiTrainOrderDetailRSV2(v *OpenApiTrainOrderDetailRSV2) {
	v.ChangeTicketInfoList = v.ChangeTicketInfoList[:0]
	v.RefundTicketInfoList = v.RefundTicketInfoList[:0]
	v.PassengerInfoList = v.PassengerInfoList[:0]
	v.PriceInfoList = v.PriceInfoList[:0]
	v.OrderBaseInfo = nil
	v.InvoiceInfo = nil
	v.TrainOrderInfo = nil
	poolOpenApiTrainOrderDetailRSV2.Put(v)
}
