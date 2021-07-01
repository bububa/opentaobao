package wdk

// PdStockDto 结构体
type PdStockDto struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 盘点类型，可选值：1：大盘  2：周盘 ；
	PdType int64 `json:"pd_type,omitempty" xml:"pd_type,omitempty"`
	// 单据号
	PdOrderCode string `json:"pd_order_code,omitempty" xml:"pd_order_code,omitempty"`
	// 店仓code，指的是盘点对象，对应一个物理店或仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 唯一识别码
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// itemList
	ItemList []PdStockDetailDto `json:"item_list,omitempty" xml:"item_list>pd_stock_detail_dto,omitempty"`
}
