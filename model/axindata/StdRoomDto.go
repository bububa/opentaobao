package axindata

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// StdRoomDto 结构体
type StdRoomDto struct {
	// 床型 BedGroupDTO 之间是 或的关系，group 内部的 BedInfoDTO 是 与的关系
	BedGroupDTOList []BedGroupDto `json:"bed_group_d_t_o_list,omitempty" xml:"bed_group_d_t_o_list>bed_group_dto,omitempty"`
	// 房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标准房型英文名称
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 宽带服务
	NetworkService string `json:"network_service,omitempty" xml:"network_service,omitempty"`
	// 设施
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 窗型0-无窗1-有窗2-部分有窗
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 卫生间信息
	Bathroom string `json:"bathroom,omitempty" xml:"bathroom,omitempty"`
	// 床型描述
	BedTypeDesc string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
	// 标准房型ID
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// 标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 是否可加床 0-否；1-是
	AddBed *model.File `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
	// 状态0 - 正常 ，-1 删除 ，-4 失效
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
}

var poolStdRoomDto = sync.Pool{
	New: func() any {
		return new(StdRoomDto)
	},
}

// GetStdRoomDto() 从对象池中获取StdRoomDto
func GetStdRoomDto() *StdRoomDto {
	return poolStdRoomDto.Get().(*StdRoomDto)
}

// ReleaseStdRoomDto 释放StdRoomDto
func ReleaseStdRoomDto(v *StdRoomDto) {
	v.BedGroupDTOList = v.BedGroupDTOList[:0]
	v.Name = ""
	v.NameEn = ""
	v.Floor = ""
	v.NetworkService = ""
	v.Facility = ""
	v.Area = ""
	v.WindowType = ""
	v.Bathroom = ""
	v.BedTypeDesc = ""
	v.Srid = 0
	v.Shid = 0
	v.MaxOccupancy = 0
	v.AddBed = nil
	v.Status = nil
	poolStdRoomDto.Put(v)
}
