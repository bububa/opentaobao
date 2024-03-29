package eleenterpriseordernew

import (
	"sync"
)

// RestaurantInfo 结构体
type RestaurantInfo struct {
	// 餐厅电话
	RestaurantTel string `json:"restaurant_tel,omitempty" xml:"restaurant_tel,omitempty"`
	// 餐厅地址
	RestaurantAddress string `json:"restaurant_address,omitempty" xml:"restaurant_address,omitempty"`
	// 餐厅名称
	RestaurantName string `json:"restaurant_name,omitempty" xml:"restaurant_name,omitempty"`
	// 餐厅唯一码
	OnlyRestaurantCode string `json:"only_restaurant_code,omitempty" xml:"only_restaurant_code,omitempty"`
	// 餐厅ID
	ErestaurantId string `json:"erestaurant_id,omitempty" xml:"erestaurant_id,omitempty"`
}

var poolRestaurantInfo = sync.Pool{
	New: func() any {
		return new(RestaurantInfo)
	},
}

// GetRestaurantInfo() 从对象池中获取RestaurantInfo
func GetRestaurantInfo() *RestaurantInfo {
	return poolRestaurantInfo.Get().(*RestaurantInfo)
}

// ReleaseRestaurantInfo 释放RestaurantInfo
func ReleaseRestaurantInfo(v *RestaurantInfo) {
	v.RestaurantTel = ""
	v.RestaurantAddress = ""
	v.RestaurantName = ""
	v.OnlyRestaurantCode = ""
	v.ErestaurantId = ""
	poolRestaurantInfo.Put(v)
}
