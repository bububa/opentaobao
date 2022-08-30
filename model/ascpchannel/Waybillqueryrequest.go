package ascpchannel

// Waybillqueryrequest 结构体
type Waybillqueryrequest struct {
	// 包裹列表
	Packages []WaybillQueryRequestData `json:"packages,omitempty" xml:"packages>waybill_query_request_data,omitempty"`
	// 发货LP单号列表
	ConsignLpOrderCodes []Consignlpordercodes `json:"consign_lp_order_codes,omitempty" xml:"consign_lp_order_codes>consignlpordercodes,omitempty"`
	// 操作人id
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 物流服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 发货LP单号
	ConsignLpOrderCode string `json:"consign_lp_order_code,omitempty" xml:"consign_lp_order_code,omitempty"`
	// 自营接口配业务模式，默认为1代表商家仓自营配 (为1时会强制校验发货的配CP和单号必须与取号时一致，且多包裹必须一次性发货)
	BusinessModel string `json:"business_model,omitempty" xml:"business_model,omitempty"`
	// 配送公司Code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}
