package qimen

// TaobaoQimenWarehouseinfoSynchronizeResponse 结构体
type TaobaoQimenWarehouseinfoSynchronizeResponse struct {
	// 仓库信息
	WarehouseInfos []WarehouseInfo `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_info,omitempty"`
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
