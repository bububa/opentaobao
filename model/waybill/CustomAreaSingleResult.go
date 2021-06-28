package waybill

// CustomAreaSingleResult 
/* model for simplify = false
type CustomAreaSingleResult struct {

    // 自定义区id
    
    CustomAreaId   int64 `json:"custom_area_id,omitempty"`
    

    // 自定义区内容的URL
    
    CustomAreaUrl   string `json:"custom_area_url,omitempty"`
    

    // keys
    
    Keys  struct {
        KeyResult  []KeyResult `json:"key_result,omitempty"`
    } `json:"keys,omitempty"`
    

}
*/

// CustomAreaSingleResult 
type CustomAreaSingleResult struct {

    // 自定义区id
    CustomAreaId   int64 `json:"custom_area_id,omitempty"`

    // 自定义区内容的URL
    CustomAreaUrl   string `json:"custom_area_url,omitempty"`

    // keys
    Keys   []KeyResult `json:"keys,omitempty"`

}
