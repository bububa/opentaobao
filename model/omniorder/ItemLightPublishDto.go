package omniorder

import (
	"sync"
)

// ItemLightPublishDto 结构体
type ItemLightPublishDto struct {
	// images
	Images []ItemLightPublishImageDto `json:"images,omitempty" xml:"images>item_light_publish_image_dto,omitempty"`
	// skus
	Skus []ItemLightPublishSkuDto `json:"skus,omitempty" xml:"skus>item_light_publish_sku_dto,omitempty"`
	// 商品条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// extendAttr
	ExtendAttr string `json:"extend_attr,omitempty" xml:"extend_attr,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 销售价
	Pretium string `json:"pretium,omitempty" xml:"pretium,omitempty"`
	// 吊牌价
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 副标题
	Subtitle string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 叶子类目ID
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 卖家ID
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolItemLightPublishDto = sync.Pool{
	New: func() any {
		return new(ItemLightPublishDto)
	},
}

// GetItemLightPublishDto() 从对象池中获取ItemLightPublishDto
func GetItemLightPublishDto() *ItemLightPublishDto {
	return poolItemLightPublishDto.Get().(*ItemLightPublishDto)
}

// ReleaseItemLightPublishDto 释放ItemLightPublishDto
func ReleaseItemLightPublishDto(v *ItemLightPublishDto) {
	v.Images = v.Images[:0]
	v.Skus = v.Skus[:0]
	v.Barcode = ""
	v.Desc = ""
	v.ExtendAttr = ""
	v.OuterId = ""
	v.Pretium = ""
	v.Price = ""
	v.Subtitle = ""
	v.Title = ""
	v.Operator = ""
	v.CatId = 0
	v.ItemId = 0
	v.UserId = 0
	poolItemLightPublishDto.Put(v)
}
