package btrip

import (
	"sync"
)

// BtriphomeResult 结构体
type BtriphomeResult struct {
	// 返回值对象
	CostCenterList []OpenCostCenterQueryRs `json:"cost_center_list,omitempty" xml:"cost_center_list>open_cost_center_query_rs,omitempty"`
	// module
	InvoiceList []InvoiceList `json:"invoice_list,omitempty" xml:"invoice_list>invoice_list,omitempty"`
	// 订单列表
	VehicleOrderList []OpenVehicleOrderRs `json:"vehicle_order_list,omitempty" xml:"vehicle_order_list>open_vehicle_order_rs,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果对象
	Module *OpenApiNewApplyRs `json:"module,omitempty" xml:"module,omitempty"`
	// module
	Invoice *OpenInvoiceDo `json:"invoice,omitempty" xml:"invoice,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBtriphomeResult = sync.Pool{
	New: func() any {
		return new(BtriphomeResult)
	},
}

// GetBtriphomeResult() 从对象池中获取BtriphomeResult
func GetBtriphomeResult() *BtriphomeResult {
	return poolBtriphomeResult.Get().(*BtriphomeResult)
}

// ReleaseBtriphomeResult 释放BtriphomeResult
func ReleaseBtriphomeResult(v *BtriphomeResult) {
	v.CostCenterList = v.CostCenterList[:0]
	v.InvoiceList = v.InvoiceList[:0]
	v.VehicleOrderList = v.VehicleOrderList[:0]
	v.ResultMsg = ""
	v.ResultCode = 0
	v.Module = nil
	v.Invoice = nil
	v.Success = false
	poolBtriphomeResult.Put(v)
}
