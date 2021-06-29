package alitripmerchant

// RoomDetailDTO 
type RoomDetailDTO struct {
    // 面积
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 最低价单位分
    LowestPrice   int64 `json:"lowest_price,omitempty" xml:"lowest_price,omitempty"`
    // 最低价单位元
    LowestPriceToString   string `json:"lowest_price_to_string,omitempty" xml:"lowest_price_to_string,omitempty"`
    // 床型详情
    RoomTypeBedInfo   *RoomTypeBedInfoDTO `json:"room_type_bed_info,omitempty" xml:"room_type_bed_info,omitempty"`
    // 入住人数图标
    MaxOccupancyIcon   string `json:"max_occupancy_icon,omitempty" xml:"max_occupancy_icon,omitempty"`
    // 货币类型
    PriceType   string `json:"price_type,omitempty" xml:"price_type,omitempty"`
    // 价格信息
    PriceInfos   []PriceInfoDTO `json:"price_infos,omitempty" xml:"price_infos>price_info_dto,omitempty"`
    // 面积图标
    AreaIcon   string `json:"area_icon,omitempty" xml:"area_icon,omitempty"`
    // 房型名称
    RoomName   string `json:"room_name,omitempty" xml:"room_name,omitempty"`
    // 房型id
    RoomId   int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
    // 宽带服务
    NetworkService   string `json:"network_service,omitempty" xml:"network_service,omitempty"`
    // 床型图标
    BedTypeIcon   string `json:"bed_type_icon,omitempty" xml:"bed_type_icon,omitempty"`
    // 是否可加床
    AddBed   int64 `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
    // 最大入住人数
    MaxOccupancy   int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
    // 房型图片
    Pics   []string `json:"pics,omitempty" xml:"pics>string,omitempty"`
    // 房型设施
    Facilitys   []RoomPropertiesDTO `json:"facilitys,omitempty" xml:"facilitys>room_properties_dto,omitempty"`
    // 是否满房
    Full   bool `json:"full,omitempty" xml:"full,omitempty"`
    // 楼层
    Floor   string `json:"floor,omitempty" xml:"floor,omitempty"`
    // 窗型
    WindowType   string `json:"window_type,omitempty" xml:"window_type,omitempty"`
    // 酒店设施
    RoomFacilities   []RoomPropertiesDTO `json:"room_facilities,omitempty" xml:"room_facilities>room_properties_dto,omitempty"`
    // 最大人数
    MaxCheckInNumber   int64 `json:"max_check_in_number,omitempty" xml:"max_check_in_number,omitempty"`
    // 床型描述
    BedTypeDesc   string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
    // 房间大小
    RoomArea   string `json:"room_area,omitempty" xml:"room_area,omitempty"`
    // 酒店名称
    HotelName   string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
    // 房间图片列表
    PictureList   []string `json:"picture_list,omitempty" xml:"picture_list>string,omitempty"`
}
