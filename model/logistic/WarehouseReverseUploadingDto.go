package logistic

import (
	"sync"
)

// WarehouseReverseUploadingDto 结构体
type WarehouseReverseUploadingDto struct {
	// 销退单列表
	UploadingReverseDTOList []UploadingReverseDto `json:"uploading_reverse_d_t_o_list,omitempty" xml:"uploading_reverse_d_t_o_list>uploading_reverse_dto,omitempty"`
}

var poolWarehouseReverseUploadingDto = sync.Pool{
	New: func() any {
		return new(WarehouseReverseUploadingDto)
	},
}

// GetWarehouseReverseUploadingDto() 从对象池中获取WarehouseReverseUploadingDto
func GetWarehouseReverseUploadingDto() *WarehouseReverseUploadingDto {
	return poolWarehouseReverseUploadingDto.Get().(*WarehouseReverseUploadingDto)
}

// ReleaseWarehouseReverseUploadingDto 释放WarehouseReverseUploadingDto
func ReleaseWarehouseReverseUploadingDto(v *WarehouseReverseUploadingDto) {
	v.UploadingReverseDTOList = v.UploadingReverseDTOList[:0]
	poolWarehouseReverseUploadingDto.Put(v)
}
