package btrip

import (
	"sync"
)

// TrainRefundTicketInfo 结构体
type TrainRefundTicketInfo struct {
	// 退票票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 乘客ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 退票金额
	RefundFee float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 服务费退款
	RefundServiceFee float64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
}

var poolTrainRefundTicketInfo = sync.Pool{
	New: func() any {
		return new(TrainRefundTicketInfo)
	},
}

// GetTrainRefundTicketInfo() 从对象池中获取TrainRefundTicketInfo
func GetTrainRefundTicketInfo() *TrainRefundTicketInfo {
	return poolTrainRefundTicketInfo.Get().(*TrainRefundTicketInfo)
}

// ReleaseTrainRefundTicketInfo 释放TrainRefundTicketInfo
func ReleaseTrainRefundTicketInfo(v *TrainRefundTicketInfo) {
	v.TicketNo = ""
	v.GmtModify = ""
	v.GmtCreate = ""
	v.UserId = ""
	v.RefundFee = 0
	v.RefundServiceFee = 0
	poolTrainRefundTicketInfo.Put(v)
}
