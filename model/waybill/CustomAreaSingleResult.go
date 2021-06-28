package waybill

// CustomAreaSingleResult 
type CustomAreaSingleResult struct {

    // 自定义区id
    
    CustomAreaId   int64 `json:"custom_area_id,omitempty" xml:"custom_area_id,omitempty"`
    

    // 自定义区内容的URL
    
    CustomAreaUrl   string `json:"custom_area_url,omitempty" xml:"custom_area_url,omitempty"`
    

    // keys
    
    Keys   []KeyResult `json:"keys,omitempty" xml:"keys,omitempty"`
    

}
