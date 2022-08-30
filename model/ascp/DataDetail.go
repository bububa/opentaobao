package ascp

// DataDetail 结构体
type DataDetail struct {
	// 详情
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// ERP货品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 响应码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 响应信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 有SKU情况下必填；无SKU情况下为空（同一个itemid不可以同时既传有sku又传空的情况）
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 仓库编码，string（50)    卖家下唯一主键
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
}
