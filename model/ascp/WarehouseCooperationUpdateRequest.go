package ascp

// WarehouseCooperationUpdateRequest 结构体
type WarehouseCooperationUpdateRequest struct {
	// 是否自动同步货主关联的所有店铺（仅菜鸟开放），为否，必填；最多50条
	SellerIds []string `json:"seller_ids,omitempty" xml:"seller_ids>string,omitempty"`
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商自定义的仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// wms奇门货主id，官方物流对接必填
	CustomerId string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
	// 合作业务： OFFICIAL_LOGISTICS（官方物流）， TIMES_AGENCY（时效代运营）， MAOCHAO_YUNCANG（猫超云仓）
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 是否自动同步货主关联的所有店铺（仅菜鸟开放） 0：否（默认） 1：是
	IsAutoSync int64 `json:"is_auto_sync,omitempty" xml:"is_auto_sync,omitempty"`
	// 合作关系状态： 1：解除合作 2：建立合作
	CooperationStatus int64 `json:"cooperation_status,omitempty" xml:"cooperation_status,omitempty"`
}
