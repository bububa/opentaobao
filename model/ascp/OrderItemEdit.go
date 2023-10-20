package ascp

import (
	"sync"
)

// OrderItemEdit 结构体
type OrderItemEdit struct {
	// 体积，单位mm3
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 叶子类目ID
	LeafCatId string `json:"leaf_cat_id,omitempty" xml:"leaf_cat_id,omitempty"`
	// 货品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 重量
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 货品类型( soft：软体，plate：版式)
	Property string `json:"property,omitempty" xml:"property,omitempty"`
	// 安装类型( whole_install：整装， part_install：组装)
	AssembleType string `json:"assemble_type,omitempty" xml:"assemble_type,omitempty"`
	// 货品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 包裹数量
	PackNum int64 `json:"pack_num,omitempty" xml:"pack_num,omitempty"`
}

var poolOrderItemEdit = sync.Pool{
	New: func() any {
		return new(OrderItemEdit)
	},
}

// GetOrderItemEdit() 从对象池中获取OrderItemEdit
func GetOrderItemEdit() *OrderItemEdit {
	return poolOrderItemEdit.Get().(*OrderItemEdit)
}

// ReleaseOrderItemEdit 释放OrderItemEdit
func ReleaseOrderItemEdit(v *OrderItemEdit) {
	v.Volume = ""
	v.LeafCatId = ""
	v.ItemName = ""
	v.Feature = ""
	v.Weight = ""
	v.Property = ""
	v.AssembleType = ""
	v.Quantity = 0
	v.PackNum = 0
	poolOrderItemEdit.Put(v)
}
