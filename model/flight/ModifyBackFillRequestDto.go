package flight

import (
	"sync"
)

// ModifyBackFillRequestDto 结构体
type ModifyBackFillRequestDto struct {
	// 改签数据
	ChangeList []ModifyItemDto `json:"change_list,omitempty" xml:"change_list>modify_item_dto,omitempty"`
	// 申请单号
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 国际国内标识:1:国内,2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolModifyBackFillRequestDto = sync.Pool{
	New: func() any {
		return new(ModifyBackFillRequestDto)
	},
}

// GetModifyBackFillRequestDto() 从对象池中获取ModifyBackFillRequestDto
func GetModifyBackFillRequestDto() *ModifyBackFillRequestDto {
	return poolModifyBackFillRequestDto.Get().(*ModifyBackFillRequestDto)
}

// ReleaseModifyBackFillRequestDto 释放ModifyBackFillRequestDto
func ReleaseModifyBackFillRequestDto(v *ModifyBackFillRequestDto) {
	v.ChangeList = v.ChangeList[:0]
	v.ApplyId = ""
	v.Currency = ""
	v.DomesticIntl = 0
	poolModifyBackFillRequestDto.Put(v)
}
