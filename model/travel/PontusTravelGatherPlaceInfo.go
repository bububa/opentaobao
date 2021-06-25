package travel

// PontusTravelGatherPlaceInfo 
type PontusTravelGatherPlaceInfo struct {

    // 集合地经纬度，英文逗号分隔。经度在前，纬度在后，英文逗号分隔最多支持6位小数，如120.111222,30.111222
    Poi   string `json:"poi,omitempty"`

    // POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
    PoiResource   string `json:"poi_resource,omitempty"`

    // 地点名称
    Name   string `json:"name,omitempty"`

}
