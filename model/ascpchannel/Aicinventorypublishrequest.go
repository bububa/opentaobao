package ascpchannel

// Aicinventorypublishrequest 结构体
type Aicinventorypublishrequest struct {
	// 库存发布请求参数
	InventoryMainOperation []Inventorymainoperation `json:"inventory_main_operation,omitempty" xml:"inventory_main_operation>inventorymainoperation,omitempty"`
}
