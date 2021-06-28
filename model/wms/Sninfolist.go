package wms

// Sninfolist 
/* model for simplify = false
type Sninfolist struct {

    // SN信息
    
    SnInfo  *struct {
        Sninfo  *Sninfo `json:"sninfo,omitempty"`
    } `json:"sn_info,omitempty"`
    

}
*/

// Sninfolist 
type Sninfolist struct {

    // SN信息
    SnInfo   *Sninfo `json:"sn_info,omitempty"`

}
