package tmallcar

import (
	"sync"
)

// Item4IsvDto 结构体
type Item4IsvDto struct {
	// 附加属性
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 商品缩略图
	ItemHeadImg string `json:"item_head_img,omitempty" xml:"item_head_img,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品服务价格,单位：元
	ServicePrice string `json:"service_price,omitempty" xml:"service_price,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItem4IsvDto = sync.Pool{
	New: func() any {
		return new(Item4IsvDto)
	},
}

// GetItem4IsvDto() 从对象池中获取Item4IsvDto
func GetItem4IsvDto() *Item4IsvDto {
	return poolItem4IsvDto.Get().(*Item4IsvDto)
}

// ReleaseItem4IsvDto 释放Item4IsvDto
func ReleaseItem4IsvDto(v *Item4IsvDto) {
	v.Extension = ""
	v.ItemHeadImg = ""
	v.ItemName = ""
	v.ServicePrice = ""
	v.ItemId = 0
	poolItem4IsvDto.Put(v)
}
