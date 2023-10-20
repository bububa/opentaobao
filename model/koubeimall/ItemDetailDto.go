package koubeimall

import (
	"sync"
)

// ItemDetailDto 结构体
type ItemDetailDto struct {
	// 详情内容
	ItemGroupDetailList []ItemGroupContent `json:"item_group_detail_list,omitempty" xml:"item_group_detail_list>item_group_content,omitempty"`
	// 图文详情
	ItemImageDetailList []ItemImageDetail `json:"item_image_detail_list,omitempty" xml:"item_image_detail_list>item_image_detail,omitempty"`
	// 补充说明
	MemoList []string `json:"memo_list,omitempty" xml:"memo_list>string,omitempty"`
	// 商品使用规则
	ItemRule *ItemRule `json:"item_rule,omitempty" xml:"item_rule,omitempty"`
	// 购买须知
	ItemBuyNotes *ItemBuyNotes `json:"item_buy_notes,omitempty" xml:"item_buy_notes,omitempty"`
	// 商品服务
	ItemServices *ItemServices `json:"item_services,omitempty" xml:"item_services,omitempty"`
}

var poolItemDetailDto = sync.Pool{
	New: func() any {
		return new(ItemDetailDto)
	},
}

// GetItemDetailDto() 从对象池中获取ItemDetailDto
func GetItemDetailDto() *ItemDetailDto {
	return poolItemDetailDto.Get().(*ItemDetailDto)
}

// ReleaseItemDetailDto 释放ItemDetailDto
func ReleaseItemDetailDto(v *ItemDetailDto) {
	v.ItemGroupDetailList = v.ItemGroupDetailList[:0]
	v.ItemImageDetailList = v.ItemImageDetailList[:0]
	v.MemoList = v.MemoList[:0]
	v.ItemRule = nil
	v.ItemBuyNotes = nil
	v.ItemServices = nil
	poolItemDetailDto.Put(v)
}
