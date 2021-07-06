package qimen

// TaobaoQimenWarehouseinfoQueryResponse 结构体
type TaobaoQimenWarehouseinfoQueryResponse struct {
	// 奇门仓储字段
	WarehouseInfos []WarehouseInfo `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_info,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	OwnerName string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
}
