package drugtrace

// WarehouseInfoDto 结构体
type WarehouseInfoDto struct {
	// 仓库位置
	WarehouseLocation string `json:"warehouse_location,omitempty" xml:"warehouse_location,omitempty"`
	// 入库日期yyyy-MM-dd
	InboundDate string `json:"inbound_date,omitempty" xml:"inbound_date,omitempty"`
	// 入库数量
	InboundSum string `json:"inbound_sum,omitempty" xml:"inbound_sum,omitempty"`
	// 保管条件
	StorageConditions string `json:"storage_conditions,omitempty" xml:"storage_conditions,omitempty"`
	// 养护周期
	MaintenanceCycle string `json:"maintenance_cycle,omitempty" xml:"maintenance_cycle,omitempty"`
	// 存储管理图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
	StoragePictures []string `json:"storage_pictures,omitempty" xml:"storage_pictures>string,omitempty"`
	// 出库数量
	OutboundSum string `json:"outbound_sum,omitempty" xml:"outbound_sum,omitempty"`
	// 出库日期yyyy-MM-dd
	OutboundDate string `json:"outbound_date,omitempty" xml:"outbound_date,omitempty"`
}
