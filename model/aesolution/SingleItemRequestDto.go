package aesolution

import (
	"sync"
)

// SingleItemRequestDto 结构体
type SingleItemRequestDto struct {
	// Content of each item, which follows different format according to different feed type.
	ItemContent string `json:"item_content,omitempty" xml:"item_content,omitempty"`
	// The id of the item_content, which could be defined by the seller. item_content_id should be unique among all the items in item_list.This field also appears in the API:aliexpress.solution.feed.query, which is regarding the convenience for the sellers to match the item_execuation_result with the item_content.
	ItemContentId string `json:"item_content_id,omitempty" xml:"item_content_id,omitempty"`
}

var poolSingleItemRequestDto = sync.Pool{
	New: func() any {
		return new(SingleItemRequestDto)
	},
}

// GetSingleItemRequestDto() 从对象池中获取SingleItemRequestDto
func GetSingleItemRequestDto() *SingleItemRequestDto {
	return poolSingleItemRequestDto.Get().(*SingleItemRequestDto)
}

// ReleaseSingleItemRequestDto 释放SingleItemRequestDto
func ReleaseSingleItemRequestDto(v *SingleItemRequestDto) {
	v.ItemContent = ""
	v.ItemContentId = ""
	poolSingleItemRequestDto.Put(v)
}
