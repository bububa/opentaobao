package nrt

import (
	"sync"
)

// ModifyFundsTypeReqDto 结构体
type ModifyFundsTypeReqDto struct {
	// 分账类型（0卖场，1摊位）
	FundsType int64 `json:"funds_type,omitempty" xml:"funds_type,omitempty"`
	// 阿里摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolModifyFundsTypeReqDto = sync.Pool{
	New: func() any {
		return new(ModifyFundsTypeReqDto)
	},
}

// GetModifyFundsTypeReqDto() 从对象池中获取ModifyFundsTypeReqDto
func GetModifyFundsTypeReqDto() *ModifyFundsTypeReqDto {
	return poolModifyFundsTypeReqDto.Get().(*ModifyFundsTypeReqDto)
}

// ReleaseModifyFundsTypeReqDto 释放ModifyFundsTypeReqDto
func ReleaseModifyFundsTypeReqDto(v *ModifyFundsTypeReqDto) {
	v.FundsType = 0
	v.StoreId = 0
	poolModifyFundsTypeReqDto.Put(v)
}
