package btrip

// HotelDetailRoomDTO 
type HotelDetailRoomDTO struct {
    // 面积
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 床型
    BedTypeString   string `json:"bed_type_string,omitempty" xml:"bed_type_string,omitempty"`
    // 设施
    Facility   string `json:"facility,omitempty" xml:"facility,omitempty"`
    // 楼层
    Floor   string `json:"floor,omitempty" xml:"floor,omitempty"`
    // 最大入住人数
    MaxOccupancy   int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
    // 房间名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 宽带服务，"0","有线上网(免费),"1","有线上网(无)","2","有线上网(收费)","3","有线上网(部分有且免费)","4","有线上网(部分有且收费)"
    NetworkService   string `json:"network_service,omitempty" xml:"network_service,omitempty"`
    // 图片
    Pics   string `json:"pics,omitempty" xml:"pics,omitempty"`
    // 详情报价列表
    Rates   []HotelDetailRateDTO `json:"rates,omitempty" xml:"rates>hotel_detail_rate_dto,omitempty"`
    // 房型ID
    Srid   int64 `json:"srid,omitempty" xml:"srid,omitempty"`
    // 状态，状态0:正常;-1:删除
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 窗型
    WindowType   string `json:"window_type,omitempty" xml:"window_type,omitempty"`
}
