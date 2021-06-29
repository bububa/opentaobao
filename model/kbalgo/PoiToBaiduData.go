package kbalgo

// PoiToBaiduData 
type PoiToBaiduData struct {

    // poiid
    
    PoiId   string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
    

    // poi明细
    
    Content   *Content `json:"content,omitempty" xml:"content,omitempty"`
    

    // 数据日期
    
    Dt   string `json:"dt,omitempty" xml:"dt,omitempty"`
    

}
