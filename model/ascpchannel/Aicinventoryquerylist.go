package ascpchannel

// Aicinventoryquerylist 结构体
type Aicinventoryquerylist struct {
	// 仓列编码列表
	StoreCodeList []string `json:"store_code_list,omitempty" xml:"store_code_list>string,omitempty"`
	// 协议ID列表
	PublishOrderNos []string `json:"publish_order_nos,omitempty" xml:"publish_order_nos>string,omitempty"`
	// 销售渠道列表
	ChannelCodeList []string `json:"channel_code_list,omitempty" xml:"channel_code_list>string,omitempty"`
	// 库存类型列表
	InventoryTypeList []int64 `json:"inventory_type_list,omitempty" xml:"inventory_type_list>int64,omitempty"`
	// 货主ID
	SourceUserId string `json:"source_user_id,omitempty" xml:"source_user_id,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
