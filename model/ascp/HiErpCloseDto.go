package ascp

// HiErpCloseDto 结构体
type HiErpCloseDto struct {
	// 履约单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 关单原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 货主id
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
}
