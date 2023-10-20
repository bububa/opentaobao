package aesolution

import (
	"sync"
)

// SingleItemResponseDto 结构体
type SingleItemResponseDto struct {
	// Execution result of each item
	ItemExecutionResult string `json:"item_execution_result,omitempty" xml:"item_execution_result,omitempty"`
	// Corresponding to the item_content_id defined by the seller when invoking the API: aliexpress.solution.feed.submit
	ItemContentId string `json:"item_content_id,omitempty" xml:"item_content_id,omitempty"`
}

var poolSingleItemResponseDto = sync.Pool{
	New: func() any {
		return new(SingleItemResponseDto)
	},
}

// GetSingleItemResponseDto() 从对象池中获取SingleItemResponseDto
func GetSingleItemResponseDto() *SingleItemResponseDto {
	return poolSingleItemResponseDto.Get().(*SingleItemResponseDto)
}

// ReleaseSingleItemResponseDto 释放SingleItemResponseDto
func ReleaseSingleItemResponseDto(v *SingleItemResponseDto) {
	v.ItemExecutionResult = ""
	v.ItemContentId = ""
	poolSingleItemResponseDto.Put(v)
}
