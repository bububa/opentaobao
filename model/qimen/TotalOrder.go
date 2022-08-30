package qimen

// TotalOrder 结构体
type TotalOrder struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 实收数量
	ActualQty string `json:"actualQty,omitempty" xml:"actualQty,omitempty"`
	// 商品名称
	ItemName string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// 仓储系统商品ID
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 商品编码
	ItemCode string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 单据行号
	OrderLineNo string `json:"orderLineNo,omitempty" xml:"orderLineNo,omitempty"`
}
