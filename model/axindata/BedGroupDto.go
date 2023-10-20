package axindata

import (
	"sync"
)

// BedGroupDto 结构体
type BedGroupDto struct {
	// 床型详细信息
	BedInfoDTOList []BedInfoDto `json:"bed_info_d_t_o_list,omitempty" xml:"bed_info_d_t_o_list>bed_info_dto,omitempty"`
}

var poolBedGroupDto = sync.Pool{
	New: func() any {
		return new(BedGroupDto)
	},
}

// GetBedGroupDto() 从对象池中获取BedGroupDto
func GetBedGroupDto() *BedGroupDto {
	return poolBedGroupDto.Get().(*BedGroupDto)
}

// ReleaseBedGroupDto 释放BedGroupDto
func ReleaseBedGroupDto(v *BedGroupDto) {
	v.BedInfoDTOList = v.BedInfoDTOList[:0]
	poolBedGroupDto.Put(v)
}
