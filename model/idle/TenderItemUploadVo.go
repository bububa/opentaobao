package idle

import (
	"sync"
)

// TenderItemUploadVo 结构体
type TenderItemUploadVo struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 类目
	Category string `json:"category,omitempty" xml:"category,omitempty"`
}

var poolTenderItemUploadVo = sync.Pool{
	New: func() any {
		return new(TenderItemUploadVo)
	},
}

// GetTenderItemUploadVo() 从对象池中获取TenderItemUploadVo
func GetTenderItemUploadVo() *TenderItemUploadVo {
	return poolTenderItemUploadVo.Get().(*TenderItemUploadVo)
}

// ReleaseTenderItemUploadVo 释放TenderItemUploadVo
func ReleaseTenderItemUploadVo(v *TenderItemUploadVo) {
	v.ItemId = ""
	v.OrderId = ""
	v.Category = ""
	poolTenderItemUploadVo.Put(v)
}
