package tuanhotel

// RelatedPoiDetailVo 结构体
type RelatedPoiDetailVo struct {
	// POI名称
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// POI ID
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}
