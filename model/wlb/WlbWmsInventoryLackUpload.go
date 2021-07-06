package wlb

// WlbWmsInventoryLackUpload 结构体
type WlbWmsInventoryLackUpload struct {
	// 商品信息列表
	ItemList []ItemListWlbWmsInventoryLackUpload `json:"item_list,omitempty" xml:"item_list>item_list_wlb_wms_inventory_lack_upload,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 外部业务编码;消息ID，用于去重
	OutBizCode string `json:"out_biz_code,omitempty" xml:"out_biz_code,omitempty"`
	// 仓储订单编码
	StoreOrderCode string `json:"store_order_code,omitempty" xml:"store_order_code,omitempty"`
	// 订单编码
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}
