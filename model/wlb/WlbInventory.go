package wlb

// WlbInventory 结构体
type WlbInventory struct {
	// 仓库编码，关联到仓库类型服务的编码非托管库存(卖家自己管理的库存，物流宝不可见又称自有库存)的所在仓库编码: STORE_SYS_PRIVATE
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// VENDIBLE--可销售库存&lt;br/&gt;FREEZE--冻结库存&lt;br/&gt;ONWAY--在途库存&lt;br/&gt;DEFECT--残次品库存
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 货主ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 可销库存数量(库存总数-拍下预扣数-占用数)
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 冻结(锁定)数量，用来跟踪库存的中间状态，比如前台销售了1件商品，这时lock加1，当商品出库的时候lock再减回去
	LockQuantity int64 `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
	// 系统自动生成
	ReserveQuantity int64 `json:"reserve_quantity,omitempty" xml:"reserve_quantity,omitempty"`
	// 系统自动生成
	OccupyQuantity int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
}
