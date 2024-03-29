package happytrip

import (
	"sync"
)

// HotelBookDto 结构体
type HotelBookDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 创建者
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 外系统订单号
	ResourceId string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 房间号
	RoomNo string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 所属订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 酒店资源id
	ResourceHotelId int64 `json:"resource_hotel_id,omitempty" xml:"resource_hotel_id,omitempty"`
	// 资源表id
	ResourceMainId int64 `json:"resource_main_id,omitempty" xml:"resource_main_id,omitempty"`
	// 出行人id
	TouristId int64 `json:"tourist_id,omitempty" xml:"tourist_id,omitempty"`
}

var poolHotelBookDto = sync.Pool{
	New: func() any {
		return new(HotelBookDto)
	},
}

// GetHotelBookDto() 从对象池中获取HotelBookDto
func GetHotelBookDto() *HotelBookDto {
	return poolHotelBookDto.Get().(*HotelBookDto)
}

// ReleaseHotelBookDto 释放HotelBookDto
func ReleaseHotelBookDto(v *HotelBookDto) {
	v.GmtCreate = ""
	v.GmtModified = ""
	v.ResourceId = ""
	v.RoomNo = ""
	v.Id = 0
	v.OrderId = 0
	v.ResourceHotelId = 0
	v.ResourceMainId = 0
	v.TouristId = 0
	poolHotelBookDto.Put(v)
}
