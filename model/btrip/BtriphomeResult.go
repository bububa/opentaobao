package btrip

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
