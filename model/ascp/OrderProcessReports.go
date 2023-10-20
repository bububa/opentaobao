package ascp

import (
	"sync"
)

// OrderProcessReports 结构体
type OrderProcessReports struct {
	// (创建发货单)条件必填
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 仓库出库，必接； 一个交易子单如果分批次发货，可以拆分多个发货单进行对接。 一个发货单如果分批次发货，分批次回传
	ConfirmOrderLines []ConfirmOrderLines `json:"confirm_order_lines,omitempty" xml:"confirm_order_lines>confirm_order_lines,omitempty"`
	// 出库包裹
	ConfirmPackages []ConfirmPackages `json:"confirm_packages,omitempty" xml:"confirm_packages>confirm_packages,omitempty"`
	// 订单标记 ，用字符串格式来表示订单标记列表
	OrderFlag string `json:"order_flag,omitempty" xml:"order_flag,omitempty"`
	// 记录行ID    随意填，请求内唯一即可，仅在返回时做结果匹配
	OrderProcessReportId string `json:"order_process_report_id,omitempty" xml:"order_process_report_id,omitempty"`
	// 单据信息
	Order *Order `json:"order,omitempty" xml:"order,omitempty"`
	// 单据作业信息
	Process *Process `json:"process,omitempty" xml:"process,omitempty"`
}

var poolOrderProcessReports = sync.Pool{
	New: func() any {
		return new(OrderProcessReports)
	},
}

// GetOrderProcessReports() 从对象池中获取OrderProcessReports
func GetOrderProcessReports() *OrderProcessReports {
	return poolOrderProcessReports.Get().(*OrderProcessReports)
}

// ReleaseOrderProcessReports 释放OrderProcessReports
func ReleaseOrderProcessReports(v *OrderProcessReports) {
	v.OrderLines = v.OrderLines[:0]
	v.ConfirmOrderLines = v.ConfirmOrderLines[:0]
	v.ConfirmPackages = v.ConfirmPackages[:0]
	v.OrderFlag = ""
	v.OrderProcessReportId = ""
	v.Order = nil
	v.Process = nil
	poolOrderProcessReports.Put(v)
}
