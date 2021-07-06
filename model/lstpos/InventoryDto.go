package lstpos

// InventoryDto 结构体
type InventoryDto struct {
	// 实时库存(库存实时结果(出参))
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// ISV商品Id
	IsvGoodsId string `json:"isv_goods_id,omitempty" xml:"isv_goods_id,omitempty"`
	// 库存删除标志(出参) ON:删除库存(库存删除后，售卖不维护库存值)   OFF:不删除库存（售卖需要维护库存值）(缺省值)
	DeleteFlag string `json:"delete_flag,omitempty" xml:"delete_flag,omitempty"`
	// 设备物理硬件ID（自身保证唯一性）
	HardwareId string `json:"hardware_id,omitempty" xml:"hardware_id,omitempty"`
	// 操作类型 STOCKTAKING-清点，盘点(set重置);STOCKENTERING-入库(加);STOCKOUTPU-出库(减);STOCK_SET-直接(set重置)
	OperateType string `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 变更库存传值为正数
	TransQuantity string `json:"trans_quantity,omitempty" xml:"trans_quantity,omitempty"`
	// isv库存Id主键
	IsvInventoryId string `json:"isv_inventory_id,omitempty" xml:"isv_inventory_id,omitempty"`
	// 设备品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 设备型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 零售通商品Id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 最后修改 精确到毫秒
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 门店对应的主账号id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 数据来源的设备类型   0：pos，1：自动售货机，-1：其它  缺省值：0
	DeviceType int64 `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 最后修改 精确到毫秒
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}
