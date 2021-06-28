package travel

// ItemScenic 
/* model for simplify = false
type ItemScenic struct {

    // 关联的套餐id
    
    RelatedPackageId   string `json:"related_package_id,omitempty"`
    

    // 必填，门票类型
    
    TicketType   string `json:"ticket_type,omitempty"`
    

    // POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
    
    PoiResource   string `json:"poi_resource,omitempty"`
    

    // 景点的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222
    
    Poi   string `json:"poi,omitempty"`
    

    // 必填，景点名称
    
    CnName   string `json:"cn_name,omitempty"`
    

    // 必填，景点所在城市
    
    City   string `json:"city,omitempty"`
    

}
*/

// ItemScenic 
type ItemScenic struct {

    // 关联的套餐id
    RelatedPackageId   string `json:"related_package_id,omitempty"`

    // 必填，门票类型
    TicketType   string `json:"ticket_type,omitempty"`

    // POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
    PoiResource   string `json:"poi_resource,omitempty"`

    // 景点的经纬度信息，经度在前，纬度在后，英文逗号分隔 最多支持6位小数，如120.111222,30.111222
    Poi   string `json:"poi,omitempty"`

    // 必填，景点名称
    CnName   string `json:"cn_name,omitempty"`

    // 必填，景点所在城市
    City   string `json:"city,omitempty"`

}
