package alitrippoi

// FliggyPoiIdParam 结构体
type FliggyPoiIdParam struct {
	// 需要查询的poiid
	PoiIds []string `json:"poi_ids,omitempty" xml:"poi_ids>string,omitempty"`
}
