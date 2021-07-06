package kbalgo

// PoiToBaiduData 结构体
type PoiToBaiduData struct {
	// poiid
	PoiId string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	// 数据日期
	Dt string `json:"dt,omitempty" xml:"dt,omitempty"`
	// poi明细
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
}
