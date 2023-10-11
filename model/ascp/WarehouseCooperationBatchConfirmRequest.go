package ascp

// WarehouseCooperationBatchConfirmRequest 结构体
type WarehouseCooperationBatchConfirmRequest struct {
	// 合作商家仓编码列表,最多100条
	CooperationWarehouses []CooperationWarehouses `json:"cooperation_warehouses,omitempty" xml:"cooperation_warehouses>cooperation_warehouses,omitempty"`
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 合作业务： ALL（全部，默认） OFFICIAL_LOGISTICS（官方物流）， TIMES_AGENCY（时效代运营）， MAOCHAO_YUNCANG（猫超云仓）
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 确认状态： 1：同意合作 2：拒绝合作
	CooperationStatus int64 `json:"cooperation_status,omitempty" xml:"cooperation_status,omitempty"`
}
