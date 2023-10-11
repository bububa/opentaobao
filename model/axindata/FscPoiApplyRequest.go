package axindata

// FscPoiApplyRequest 结构体
type FscPoiApplyRequest struct {
	// POI新增列表
	PoiList []FscPoiApplyApiDto `json:"poi_list,omitempty" xml:"poi_list>fsc_poi_apply_api_dto,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
