package servicecenter

import (
	"sync"
)

// CarItemInfoDto 结构体
type CarItemInfoDto struct {
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 车系
	Line string `json:"line,omitempty" xml:"line,omitempty"`
	// 型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 不会返回
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 年款
	Year string `json:"year,omitempty" xml:"year,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 不会返回
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// skuId不会返回
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolCarItemInfoDto = sync.Pool{
	New: func() any {
		return new(CarItemInfoDto)
	},
}

// GetCarItemInfoDto() 从对象池中获取CarItemInfoDto
func GetCarItemInfoDto() *CarItemInfoDto {
	return poolCarItemInfoDto.Get().(*CarItemInfoDto)
}

// ReleaseCarItemInfoDto 释放CarItemInfoDto
func ReleaseCarItemInfoDto(v *CarItemInfoDto) {
	v.Brand = ""
	v.Line = ""
	v.Model = ""
	v.SellerNick = ""
	v.Year = ""
	v.ItemId = 0
	v.SellerId = 0
	v.SkuId = 0
	poolCarItemInfoDto.Put(v)
}
