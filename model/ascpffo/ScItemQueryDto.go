package ascpffo

import (
	"sync"
)

// ScItemQueryDto 结构体
type ScItemQueryDto struct {
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品条码
	WhcBarcode string `json:"whc_barcode,omitempty" xml:"whc_barcode,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 货品Id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolScItemQueryDto = sync.Pool{
	New: func() any {
		return new(ScItemQueryDto)
	},
}

// GetScItemQueryDto() 从对象池中获取ScItemQueryDto
func GetScItemQueryDto() *ScItemQueryDto {
	return poolScItemQueryDto.Get().(*ScItemQueryDto)
}

// ReleaseScItemQueryDto 释放ScItemQueryDto
func ReleaseScItemQueryDto(v *ScItemQueryDto) {
	v.ScItemCode = ""
	v.WhcBarcode = ""
	v.BizType = 0
	v.PageIndex = 0
	v.PageSize = 0
	v.ScItemId = 0
	poolScItemQueryDto.Put(v)
}
