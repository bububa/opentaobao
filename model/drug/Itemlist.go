package drug

import (
	"sync"
)

// Itemlist 结构体
type Itemlist struct {
	// itemId
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// isO2O
	IsO2o string `json:"is_o2o,omitempty" xml:"is_o2o,omitempty"`
	// itemName
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// listPicUrl
	ListPicUrl string `json:"list_pic_url,omitempty" xml:"list_pic_url,omitempty"`
	// oriPrice
	OriPrice string `json:"ori_price,omitempty" xml:"ori_price,omitempty"`
	// price
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// symptom
	Symptom string `json:"symptom,omitempty" xml:"symptom,omitempty"`
	// quantity
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// deliveryTime
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// deliveryTypeDesc
	DeliveryTypeDesc string `json:"delivery_type_desc,omitempty" xml:"delivery_type_desc,omitempty"`
	// backCate
	BackCate int64 `json:"back_cate,omitempty" xml:"back_cate,omitempty"`
	// atLimit
	AtLimit int64 `json:"at_limit,omitempty" xml:"at_limit,omitempty"`
	// deliveryType
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// rx
	Rx bool `json:"rx,omitempty" xml:"rx,omitempty"`
}

var poolItemlist = sync.Pool{
	New: func() any {
		return new(Itemlist)
	},
}

// GetItemlist() 从对象池中获取Itemlist
func GetItemlist() *Itemlist {
	return poolItemlist.Get().(*Itemlist)
}

// ReleaseItemlist 释放Itemlist
func ReleaseItemlist(v *Itemlist) {
	v.ItemId = ""
	v.IsO2o = ""
	v.ItemName = ""
	v.ListPicUrl = ""
	v.OriPrice = ""
	v.Price = ""
	v.Symptom = ""
	v.Quantity = ""
	v.DeliveryTime = ""
	v.DeliveryTypeDesc = ""
	v.BackCate = 0
	v.AtLimit = 0
	v.DeliveryType = 0
	v.Rx = false
	poolItemlist.Put(v)
}
