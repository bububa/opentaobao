package ascp

// WarehouseStatusUpdateRequest 结构体
type WarehouseStatusUpdateRequest struct {
	// 资源编码，string（50)    货主下唯一主键 仓资源或者配资源的唯一编码（会员ID+仓ID/配ID）
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
	// 状态：0=停用；1=启用
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}
