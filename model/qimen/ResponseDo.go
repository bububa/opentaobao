package qimen

import (
	"sync"
)

// ResponseDo 结构体
type ResponseDo struct {
	// itemInventories
	ItemInventories []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories>item_inventory,omitempty"`
	// 响应结果:success|failure,success,string(10),必填,
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码,0,string(50),,
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息,invalid appkey,string(100),,
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// item
	Item *Item `json:"item,omitempty" xml:"item,omitempty"`
}

var poolResponseDo = sync.Pool{
	New: func() any {
		return new(ResponseDo)
	},
}

// GetResponseDo() 从对象池中获取ResponseDo
func GetResponseDo() *ResponseDo {
	return poolResponseDo.Get().(*ResponseDo)
}

// ReleaseResponseDo 释放ResponseDo
func ReleaseResponseDo(v *ResponseDo) {
	v.ItemInventories = v.ItemInventories[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.Item = nil
	poolResponseDo.Put(v)
}
