package ascpchannel

// Ownerdto 结构体
type Ownerdto struct {
	// 物流货主ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 供应商ID
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
