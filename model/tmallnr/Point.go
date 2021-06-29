package tmallnr

// Point 
type Point struct {
    // 纬度
    Lat   string `json:"lat,omitempty" xml:"lat,omitempty"`
    // 经度
    Lng   string `json:"lng,omitempty" xml:"lng,omitempty"`
}
