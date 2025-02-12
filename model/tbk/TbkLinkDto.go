package tbk

import (
	"sync"
)

// TbkLinkDto 结构体
type TbkLinkDto struct {
	// 转链结果
	MaterialUrlList []MaterialUrlList `json:"material_url_list,omitempty" xml:"material_url_list>material_url_list,omitempty"`
	// 店铺转链结果
	ShopUrlList []ShopUrlList `json:"shop_url_list,omitempty" xml:"shop_url_list>shop_url_list,omitempty"`
	// 活动转链信息
	EventUrlList []EventUrlList `json:"event_url_list,omitempty" xml:"event_url_list>event_url_list,omitempty"`
	// 单品转链信息
	ItemUrlList []ItemUrlList `json:"item_url_list,omitempty" xml:"item_url_list>item_url_list,omitempty"`
}

var poolTbkLinkDto = sync.Pool{
	New: func() any {
		return new(TbkLinkDto)
	},
}

// GetTbkLinkDto() 从对象池中获取TbkLinkDto
func GetTbkLinkDto() *TbkLinkDto {
	return poolTbkLinkDto.Get().(*TbkLinkDto)
}

// ReleaseTbkLinkDto 释放TbkLinkDto
func ReleaseTbkLinkDto(v *TbkLinkDto) {
	v.MaterialUrlList = v.MaterialUrlList[:0]
	v.ShopUrlList = v.ShopUrlList[:0]
	v.EventUrlList = v.EventUrlList[:0]
	v.ItemUrlList = v.ItemUrlList[:0]
	poolTbkLinkDto.Put(v)
}
