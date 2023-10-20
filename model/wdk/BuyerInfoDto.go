package wdk

import (
	"sync"
)

// BuyerInfoDto 结构体
type BuyerInfoDto struct {
	// 买家姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// 买家电话号码
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 收货地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
}

var poolBuyerInfoDto = sync.Pool{
	New: func() any {
		return new(BuyerInfoDto)
	},
}

// GetBuyerInfoDto() 从对象池中获取BuyerInfoDto
func GetBuyerInfoDto() *BuyerInfoDto {
	return poolBuyerInfoDto.Get().(*BuyerInfoDto)
}

// ReleaseBuyerInfoDto 释放BuyerInfoDto
func ReleaseBuyerInfoDto(v *BuyerInfoDto) {
	v.BuyerName = ""
	v.BuyerPhone = ""
	v.BuyerAddress = ""
	poolBuyerInfoDto.Put(v)
}
