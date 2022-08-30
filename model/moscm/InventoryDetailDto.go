package moscm

// InventoryDetailDto 结构体
type InventoryDetailDto struct {
	// 专柜号
	CounterNo string `json:"counter_no,omitempty" xml:"counter_no,omitempty"`
	// 专柜名
	CounterName string `json:"counter_name,omitempty" xml:"counter_name,omitempty"`
	// 门店号
	StoreNo string `json:"store_no,omitempty" xml:"store_no,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// outer id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// sku id
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 颜色/尺码/规格
	SaleProperty string `json:"sale_property,omitempty" xml:"sale_property,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 款号
	StyleNo string `json:"style_no,omitempty" xml:"style_no,omitempty"`
	// 仓号
	WarehouseNumber string `json:"warehouse_number,omitempty" xml:"warehouse_number,omitempty"`
	// 仓库名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 占库数量
	OccupyQty string `json:"occupy_qty,omitempty" xml:"occupy_qty,omitempty"`
	// 在库数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 残次品数量
	DefectiveQty string `json:"defective_qty,omitempty" xml:"defective_qty,omitempty"`
	// 商品状态
	GoodsStatus string `json:"goods_status,omitempty" xml:"goods_status,omitempty"`
	// 商品类型
	GoodsType string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// 货号
	ArtNo string `json:"art_no,omitempty" xml:"art_no,omitempty"`
}
