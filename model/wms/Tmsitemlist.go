package wms

// Tmsitemlist 
/* model for simplify = false
type Tmsitemlist struct {

    // 包裹里面商品
    
    TmsItem  *struct {
        Tmsitem  *Tmsitem `json:"tmsitem,omitempty"`
    } `json:"tms_item,omitempty"`
    

}
*/

// Tmsitemlist 
type Tmsitemlist struct {

    // 包裹里面商品
    TmsItem   *Tmsitem `json:"tms_item,omitempty"`

}
