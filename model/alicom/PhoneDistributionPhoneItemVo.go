package alicom

import (
	"sync"
)

// PhoneDistributionPhoneItemVo 结构体
type PhoneDistributionPhoneItemVo struct {
	// 商品面额
	Face string `json:"face,omitempty" xml:"face,omitempty"`
	// 商品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品标识
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolPhoneDistributionPhoneItemVo = sync.Pool{
	New: func() any {
		return new(PhoneDistributionPhoneItemVo)
	},
}

// GetPhoneDistributionPhoneItemVo() 从对象池中获取PhoneDistributionPhoneItemVo
func GetPhoneDistributionPhoneItemVo() *PhoneDistributionPhoneItemVo {
	return poolPhoneDistributionPhoneItemVo.Get().(*PhoneDistributionPhoneItemVo)
}

// ReleasePhoneDistributionPhoneItemVo 释放PhoneDistributionPhoneItemVo
func ReleasePhoneDistributionPhoneItemVo(v *PhoneDistributionPhoneItemVo) {
	v.Face = ""
	v.Price = ""
	v.ItemId = 0
	poolPhoneDistributionPhoneItemVo.Put(v)
}
