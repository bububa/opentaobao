package axindata

import (
	"github.com/bububa/opentaobao/model"
)

// StdRoomType 结构体
type StdRoomType struct {
	// 床型 BedGroupDTO 之间是 或的关系，group 内部的 BedInfoDTO 是 与的关系
	BedGroupDTOList []BedGroupDto `json:"bed_group_d_t_o_list,omitempty" xml:"bed_group_d_t_o_list>bed_group_dto,omitempty"`
	// 标准房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 宽带服务（&#34;0&#34;,&#34;有线上网(免费),&#34;1&#34;,&#34;有线上网(无)&#34;,&#34;2&#34;,&#34;有线上网(收费)&#34;,&#34;3&#34;,&#34;有线上网(部分有且免费)&#34;,&#34;4&#34;,&#34;有线上网(部分有且收费)&#34;）
	NetworkService string `json:"network_service,omitempty" xml:"network_service,omitempty"`
	// 设施
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 窗型（0, 无窗,1, 有窗;）
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 床型
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 床型描述
	BedTypeDesc string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
	// 标准房型id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 状态（0 - 正常 ，-1 删除 ，-4 失效）
	Status *model.File `json:"status,omitempty" xml:"status,omitempty"`
}
