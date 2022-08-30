package alihealth2

// FutureItemError 结构体
type FutureItemError struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 错误编码, 编码以"ALIHEALTH.BC.5"开头时表示当前品仓的操作处于未知状态, 后继可通过同样的请求流水号来获取最终结果
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 期货实际已减少数量
	RealChangeQuantity int64 `json:"real_change_quantity,omitempty" xml:"real_change_quantity,omitempty"`
	// 期货理论可减少数量
	CanChangeQuantity int64 `json:"can_change_quantity,omitempty" xml:"can_change_quantity,omitempty"`
}
