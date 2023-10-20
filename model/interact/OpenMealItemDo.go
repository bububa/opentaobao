package interact

import (
	"sync"
)

// OpenMealItemDo 结构体
type OpenMealItemDo struct {
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolOpenMealItemDo = sync.Pool{
	New: func() any {
		return new(OpenMealItemDo)
	},
}

// GetOpenMealItemDo() 从对象池中获取OpenMealItemDo
func GetOpenMealItemDo() *OpenMealItemDo {
	return poolOpenMealItemDo.Get().(*OpenMealItemDo)
}

// ReleaseOpenMealItemDo 释放OpenMealItemDo
func ReleaseOpenMealItemDo(v *OpenMealItemDo) {
	v.ItemId = 0
	poolOpenMealItemDo.Put(v)
}
