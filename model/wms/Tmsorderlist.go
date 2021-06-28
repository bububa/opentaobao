package wms

// Tmsorderlist 
/* model for simplify = false
type Tmsorderlist struct {

    // 运单信息列表
    
    TmsOrder  *struct {
        Tmsorder  *Tmsorder `json:"tmsorder,omitempty"`
    } `json:"tms_order,omitempty"`
    

}
*/

// Tmsorderlist 
type Tmsorderlist struct {

    // 运单信息列表
    TmsOrder   *Tmsorder `json:"tms_order,omitempty"`

}
