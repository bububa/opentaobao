package ascp

import (
	"sync"
)

// WmsOrderProcessReportRequest 结构体
type WmsOrderProcessReportRequest struct {
	// (创建发货单)条件必填
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 仓库出库，必接； 一个交易子单如果分批次发货，可以拆分多个发货单进行对接。 一个发货单如果分批次发货，分批次回传
	ConfirmOrderLines []ConfirmOrderLines `json:"confirm_order_lines,omitempty" xml:"confirm_order_lines>confirm_order_lines,omitempty"`
	// 出库包裹
	ConfirmPackages []ConfirmPackages `json:"confirm_packages,omitempty" xml:"confirm_packages>confirm_packages,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单标记 ，用字符串格式来表示订单标记列表
	OrderFlag string `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 单据信息
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 单据作业信息
	Process *Process `json:"process,omitempty" xml:"process,omitempty"`
}

var poolWmsOrderProcessReportRequest = sync.Pool{
	New: func() any {
		return new(WmsOrderProcessReportRequest)
	},
}

// GetWmsOrderProcessReportRequest() 从对象池中获取WmsOrderProcessReportRequest
func GetWmsOrderProcessReportRequest() *WmsOrderProcessReportRequest {
	return poolWmsOrderProcessReportRequest.Get().(*WmsOrderProcessReportRequest)
}

// ReleaseWmsOrderProcessReportRequest 释放WmsOrderProcessReportRequest
func ReleaseWmsOrderProcessReportRequest(v *WmsOrderProcessReportRequest) {
	v.OrderLines = v.OrderLines[:0]
	v.ConfirmOrderLines = v.ConfirmOrderLines[:0]
	v.ConfirmPackages = v.ConfirmPackages[:0]
	v.RequestId = ""
	v.OrderFlag = ""
	v.RequestTime = 0
	v.Order = nil
	v.Process = nil
	poolWmsOrderProcessReportRequest.Put(v)
}
