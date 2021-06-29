package hotelhstdf

// SroomTypePo 
type SroomTypePo struct {
    // 标准房型id
    Srid   int64 `json:"srid,omitempty" xml:"srid,omitempty"`
    // 标准房型名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 标准房型英文名称
    NameE   string `json:"name_e,omitempty" xml:"name_e,omitempty"`
    // 标准房型床型
    BedType   string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
    // 标准房型面积
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    // 窗型，0:无窗,1:有窗,2:部分有窗,3:暗窗,4:部分暗窗
    WindowType   int64 `json:"window_type,omitempty" xml:"window_type,omitempty"`
}
