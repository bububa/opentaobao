package simba

import (
	"sync"
)

// SubwayItem 结构体
type SubwayItem struct {
	// 商品信息在外部系统（淘宝主站）的标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品信息在外部系统（淘宝主站）的价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// imgUrl
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 商品信息在外部系统(淘宝主站)的数字id
	NumId int64 `json:"num_id,omitempty" xml:"num_id,omitempty"`
	// 扩展属性对象
	ExtraAttributes *ExtraAttributes `json:"extra_attributes,omitempty" xml:"extra_attributes,omitempty"`
}

var poolSubwayItem = sync.Pool{
	New: func() any {
		return new(SubwayItem)
	},
}

// GetSubwayItem() 从对象池中获取SubwayItem
func GetSubwayItem() *SubwayItem {
	return poolSubwayItem.Get().(*SubwayItem)
}

// ReleaseSubwayItem 释放SubwayItem
func ReleaseSubwayItem(v *SubwayItem) {
	v.Title = ""
	v.Price = ""
	v.ImgUrl = ""
	v.NumId = 0
	v.ExtraAttributes = nil
	poolSubwayItem.Put(v)
}
