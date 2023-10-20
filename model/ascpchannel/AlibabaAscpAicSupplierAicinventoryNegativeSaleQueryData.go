package ascpchannel

// AlibabaascpaicsupplieraicinventorynegativesalequeryData 结构体
type AlibabaascpaicsupplieraicinventorynegativesalequeryData struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 库存总量(前台)
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 占用量
	LockQuantity string `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
	// 预留库存总量
	ReservedQuantity string `json:"reserved_quantity,omitempty" xml:"reserved_quantity,omitempty"`
	// 预留库存占用量
	ReservedLockQuantity string `json:"reserved_lock_quantity,omitempty" xml:"reserved_lock_quantity,omitempty"`
	// 协议ID=plan_id单据号
	PublishOrderNo string `json:"publish_order_no,omitempty" xml:"publish_order_no,omitempty"`
	// 当前可售库存量(前台)
	SellableQuantity string `json:"sellable_quantity,omitempty" xml:"sellable_quantity,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 物流货主id
	SourceUserId int64 `json:"source_user_id,omitempty" xml:"source_user_id,omitempty"`
	// 有效状态
	InvStatus int64 `json:"inv_status,omitempty" xml:"inv_status,omitempty"`
	// ASCP销售市场中的计划ID，供应商不能用这个
	TradeFutureInvId int64 `json:"trade_future_inv_id,omitempty" xml:"trade_future_inv_id,omitempty"`
}
