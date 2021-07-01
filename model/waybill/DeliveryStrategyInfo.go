package waybill

// DeliveryStrategyInfo 结构体
type DeliveryStrategyInfo struct {
	// 识别买家备注: 0-忽略, 1-识别, 2-仅识别合作cp
	BuyerMessageRule int64 `json:"buyer_message_rule,omitempty" xml:"buyer_message_rule,omitempty"`
	// 合作CP信息
	CocpInfoList []CpInfo `json:"cocp_info_list,omitempty" xml:"cocp_info_list>cp_info,omitempty"`
	// 特殊线路
	SpecialRouteInfoList []SpecialRouteInfo `json:"special_route_info_list,omitempty" xml:"special_route_info_list>special_route_info,omitempty"`
	// 仓id
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// 仓名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
}
