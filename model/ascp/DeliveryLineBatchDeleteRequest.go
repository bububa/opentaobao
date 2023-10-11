package ascp

// DeliveryLineBatchDeleteRequest 结构体
type DeliveryLineBatchDeleteRequest struct {
	// 删除规则的收货地（组）；当删除类型为2时必填
	Addresses []Addresse `json:"addresses,omitempty" xml:"addresses>addresse,omitempty"`
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 删除类型： 1-by仓删除所有到货规则 2-by仓+收货地删除到货规则
	DeleteType int64 `json:"delete_type,omitempty" xml:"delete_type,omitempty"`
	// 线路规则类型： 1-通用规则； 2-个性化规则
	RuleType int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
}
