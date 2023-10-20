package aesolution

import (
	"sync"
)

// ItemListResultDto 结构体
type ItemListResultDto struct {
	// product list
	AeopAEProductDisplayDTOList []ItemDisplayDto `json:"aeop_a_e_product_display_d_t_o_list,omitempty" xml:"aeop_a_e_product_display_d_t_o_list>item_display_dto,omitempty"`
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// error msg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// error code
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// total page
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// products total count
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// current page
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// success or not
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolItemListResultDto = sync.Pool{
	New: func() any {
		return new(ItemListResultDto)
	},
}

// GetItemListResultDto() 从对象池中获取ItemListResultDto
func GetItemListResultDto() *ItemListResultDto {
	return poolItemListResultDto.Get().(*ItemListResultDto)
}

// ReleaseItemListResultDto 释放ItemListResultDto
func ReleaseItemListResultDto(v *ItemListResultDto) {
	v.AeopAEProductDisplayDTOList = v.AeopAEProductDisplayDTOList[:0]
	v.ErrorMessage = ""
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.TotalPage = 0
	v.ProductCount = 0
	v.CurrentPage = 0
	v.Success = false
	poolItemListResultDto.Put(v)
}
