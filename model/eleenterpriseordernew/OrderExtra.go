package eleenterpriseordernew

import (
	"sync"
)

// OrderExtra 结构体
type OrderExtra struct {
	// 费用
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 费用项名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 订单项目分类（参考附录）
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolOrderExtra = sync.Pool{
	New: func() any {
		return new(OrderExtra)
	},
}

// GetOrderExtra() 从对象池中获取OrderExtra
func GetOrderExtra() *OrderExtra {
	return poolOrderExtra.Get().(*OrderExtra)
}

// ReleaseOrderExtra 释放OrderExtra
func ReleaseOrderExtra(v *OrderExtra) {
	v.Price = ""
	v.Name = ""
	v.Quantity = 0
	v.CategoryId = 0
	poolOrderExtra.Put(v)
}
