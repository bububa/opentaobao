package hotelhstdf

// RoomTypePo 
type RoomTypePo struct {
    // 外部系统id
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 卖家未匹配房型的面积
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 卖家未匹配房型的床型
    BedType   string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
    // 卖家未匹配房型英文名称
    NameE   string `json:"name_e,omitempty" xml:"name_e,omitempty"`
    // 卖家未匹配房型名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 卖家未匹配rid
    Rid   int64 `json:"rid,omitempty" xml:"rid,omitempty"`
    // 窗型
    WindowType   int64 `json:"window_type,omitempty" xml:"window_type,omitempty"`
}
