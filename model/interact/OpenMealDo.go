package interact

import (
	"sync"
)

// OpenMealDo 结构体
type OpenMealDo struct {
	// 套餐商品列表
	Items []OpenMealItemDo `json:"items,omitempty" xml:"items>open_meal_item_do,omitempty"`
	// 手淘购买链接
	H5BuyUrl string `json:"h5_buy_url,omitempty" xml:"h5_buy_url,omitempty"`
	// 套餐名称
	MealName string `json:"meal_name,omitempty" xml:"meal_name,omitempty"`
	// 套餐开始时间戳
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 套餐结束时间戳
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 单位分，套餐总价
	MealPrice int64 `json:"meal_price,omitempty" xml:"meal_price,omitempty"`
}

var poolOpenMealDo = sync.Pool{
	New: func() any {
		return new(OpenMealDo)
	},
}

// GetOpenMealDo() 从对象池中获取OpenMealDo
func GetOpenMealDo() *OpenMealDo {
	return poolOpenMealDo.Get().(*OpenMealDo)
}

// ReleaseOpenMealDo 释放OpenMealDo
func ReleaseOpenMealDo(v *OpenMealDo) {
	v.Items = v.Items[:0]
	v.H5BuyUrl = ""
	v.MealName = ""
	v.StartTime = ""
	v.EndTime = ""
	v.MealPrice = 0
	poolOpenMealDo.Put(v)
}
