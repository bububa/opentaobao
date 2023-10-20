package fenxiao

import (
	"sync"
)

// InventorySum 结构体
type InventorySum struct {
	// 商家仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 库存类型名称
	InventoryTypeName string `json:"inventory_type_name,omitempty" xml:"inventory_type_name,omitempty"`
	// 商品商家编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 总预扣数量
	ReserveQuantity int64 `json:"reserve_quantity,omitempty" xml:"reserve_quantity,omitempty"`
	// 库存类型：&lt;br/&gt;1：正常 &lt;br/&gt;2：损坏 &lt;br/&gt;3：冻结 &lt;br/&gt;10：质押 &lt;br/&gt;11-20:商家自定义
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 商品后端ID，如果有传sc_item_code,参数可以为0
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 总物理库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 总占用数量
	OccupyQuantity int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
}

var poolInventorySum = sync.Pool{
	New: func() any {
		return new(InventorySum)
	},
}

// GetInventorySum() 从对象池中获取InventorySum
func GetInventorySum() *InventorySum {
	return poolInventorySum.Get().(*InventorySum)
}

// ReleaseInventorySum 释放InventorySum
func ReleaseInventorySum(v *InventorySum) {
	v.StoreCode = ""
	v.InventoryTypeName = ""
	v.ScItemCode = ""
	v.ReserveQuantity = 0
	v.InventoryType = 0
	v.ScItemId = 0
	v.Quantity = 0
	v.OccupyQuantity = 0
	poolInventorySum.Put(v)
}
