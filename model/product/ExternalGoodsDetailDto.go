package product

import (
	"sync"
)

// ExternalGoodsDetailDto 结构体
type ExternalGoodsDetailDto struct {
	// 商品图片url列表
	ImageList []GoodsImageDto `json:"image_list,omitempty" xml:"image_list>goods_image_dto,omitempty"`
	// 游戏属性对
	GoodsPropertyList []GoodsPropertyValueDto `json:"goods_property_list,omitempty" xml:"goods_property_list>goods_property_value_dto,omitempty"`
	// 卖家账号信息商品属性对
	SellerAccountPropertyList []GoodsPropertyValueDto `json:"seller_account_property_list,omitempty" xml:"seller_account_property_list>goods_property_value_dto,omitempty"`
	// 商品基本信息
	GoodsBaseInfo *GoodsBaseInfoDto `json:"goods_base_info,omitempty" xml:"goods_base_info,omitempty"`
	// 类目
	Category *GoodsCategoryDto `json:"category,omitempty" xml:"category,omitempty"`
	// 游戏属性
	GameProperty *GamePropertyDto `json:"game_property,omitempty" xml:"game_property,omitempty"`
	// 是否支持找回包赔
	SupportRetrieveCompensation bool `json:"support_retrieve_compensation,omitempty" xml:"support_retrieve_compensation,omitempty"`
	// 是否支持议价
	CanBargain bool `json:"can_bargain,omitempty" xml:"can_bargain,omitempty"`
}

var poolExternalGoodsDetailDto = sync.Pool{
	New: func() any {
		return new(ExternalGoodsDetailDto)
	},
}

// GetExternalGoodsDetailDto() 从对象池中获取ExternalGoodsDetailDto
func GetExternalGoodsDetailDto() *ExternalGoodsDetailDto {
	return poolExternalGoodsDetailDto.Get().(*ExternalGoodsDetailDto)
}

// ReleaseExternalGoodsDetailDto 释放ExternalGoodsDetailDto
func ReleaseExternalGoodsDetailDto(v *ExternalGoodsDetailDto) {
	v.ImageList = v.ImageList[:0]
	v.GoodsPropertyList = v.GoodsPropertyList[:0]
	v.SellerAccountPropertyList = v.SellerAccountPropertyList[:0]
	v.GoodsBaseInfo = nil
	v.Category = nil
	v.GameProperty = nil
	v.SupportRetrieveCompensation = false
	v.CanBargain = false
	poolExternalGoodsDetailDto.Put(v)
}
