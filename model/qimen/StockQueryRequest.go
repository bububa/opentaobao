package qimen

// StockQueryRequest 结构体
type StockQueryRequest struct {
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 商品编码
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 仓储系统商品ID
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	InventoryType string `json:"inventoryType,omitempty" xml:"inventoryType,omitempty"`
	// 批次编码
	BatchCode string `json:"batchCode,omitempty" xml:"batchCode,omitempty"`
	// 商品生产日期(YYYY-MM-DD)
	ProductDate string `json:"productDate,omitempty" xml:"productDate,omitempty"`
	// 商品过期日期(YYYY-MM-DD)
	ExpireDate string `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 当前页(从1开始)
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数(最多100条)
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
