package hotelhstdf

// SroomTypeMatchParam 
type SroomTypeMatchParam struct {
    // 标准房型id
    Srid   int64 `json:"srid,omitempty" xml:"srid,omitempty"`
    // 卖家房型id
    Rid   int64 `json:"rid,omitempty" xml:"rid,omitempty"`
}
