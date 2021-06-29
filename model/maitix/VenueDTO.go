package maitix

// VenueDto 
type VenueDto struct {
    // 场馆id
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 纬度
    Lat   string `json:"lat,omitempty" xml:"lat,omitempty"`
    // 经度
    Lng   string `json:"lng,omitempty" xml:"lng,omitempty"`
    // 场馆名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 场馆地址
    VenueAddress   string `json:"venue_address,omitempty" xml:"venue_address,omitempty"`
}
