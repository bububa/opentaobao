package alihealth2

// FutureItem 结构体
type FutureItem struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 正数表示增加期货, 负数表示减少期货
	ChangeQuantity int64 `json:"change_quantity,omitempty" xml:"change_quantity,omitempty"`
	// 货品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
