package alihealth2

import (
	"sync"
)

// Ingredients 结构体
type Ingredients struct {
	// 成分单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 成分名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 成分数值
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolIngredients = sync.Pool{
	New: func() any {
		return new(Ingredients)
	},
}

// GetIngredients() 从对象池中获取Ingredients
func GetIngredients() *Ingredients {
	return poolIngredients.Get().(*Ingredients)
}

// ReleaseIngredients 释放Ingredients
func ReleaseIngredients(v *Ingredients) {
	v.Unit = ""
	v.Name = ""
	v.Quantity = 0
	poolIngredients.Put(v)
}
