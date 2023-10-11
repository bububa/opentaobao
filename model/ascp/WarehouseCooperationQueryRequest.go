package ascp

// WarehouseCooperationQueryRequest 结构体
type WarehouseCooperationQueryRequest struct {
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 合作业务： ALL（全部，默认） OFFICIAL_LOGISTICS（官方物流）， TIMES_AGENCY（时效代运营）， MAOCHAO_YUNCANG（猫超云仓）
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 合作状态（默认，全部状态查询） 1：合作待确认 2：合作中 3：已拒绝合作 4：商家取消合作 5：服务商取消合作
	CooperationStatus int64 `json:"cooperation_status,omitempty" xml:"cooperation_status,omitempty"`
	// 第几页
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每一页多少条，不超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
