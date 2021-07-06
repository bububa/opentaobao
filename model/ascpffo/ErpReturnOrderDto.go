package ascpffo

// ErpReturnOrderDto 结构体
type ErpReturnOrderDto struct {
	// 单据创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
	// 库存类型描述
	InventoryTypeDesc string `json:"inventory_type_desc,omitempty" xml:"inventory_type_desc,omitempty"`
	// 物流单号
	LbxNo string `json:"lbx_no,omitempty" xml:"lbx_no,omitempty"`
	// 退供单号
	ReturnOrderNo string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
	// 退供原因
	ReturnReasonDesc string `json:"return_reason_desc,omitempty" xml:"return_reason_desc,omitempty"`
	// 退供类型
	ReturnTypeDesc string `json:"return_type_desc,omitempty" xml:"return_type_desc,omitempty"`
	// 状态描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 仓名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 退供总金额,单位分
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 实际退供总金额，单位分
	TotalReturnAmount string `json:"total_return_amount,omitempty" xml:"total_return_amount,omitempty"`
	// 单据创建时间戳
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 出库时间戳
	GmtOutbound int64 `json:"gmt_outbound,omitempty" xml:"gmt_outbound,omitempty"`
	// 退供SKU数量
	SkuCount int64 `json:"sku_count,omitempty" xml:"sku_count,omitempty"`
	// 状态编码
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 供应商Id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 退供总数量
	TotalQuantity int64 `json:"total_quantity,omitempty" xml:"total_quantity,omitempty"`
	// 实际退供数量
	TotalReturnQuantity int64 `json:"total_return_quantity,omitempty" xml:"total_return_quantity,omitempty"`
}
