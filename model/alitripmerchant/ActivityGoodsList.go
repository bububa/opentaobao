package alitripmerchant

import (
	"sync"
)

// ActivityGoodsList 结构体
type ActivityGoodsList struct {
	// 奖品图片
	Image string `json:"image,omitempty" xml:"image,omitempty"`
	// 奖品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 奖品id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolActivityGoodsList = sync.Pool{
	New: func() any {
		return new(ActivityGoodsList)
	},
}

// GetActivityGoodsList() 从对象池中获取ActivityGoodsList
func GetActivityGoodsList() *ActivityGoodsList {
	return poolActivityGoodsList.Get().(*ActivityGoodsList)
}

// ReleaseActivityGoodsList 释放ActivityGoodsList
func ReleaseActivityGoodsList(v *ActivityGoodsList) {
	v.Image = ""
	v.Name = ""
	v.Id = 0
	poolActivityGoodsList.Put(v)
}
