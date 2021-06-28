package wms

// Invoinceconfirmlist 
/* model for simplify = false
type Invoinceconfirmlist struct {

    // 发票确认信息
    
    InvoinceConfirm  *struct {
        Invoinceconfirm  *Invoinceconfirm `json:"invoinceconfirm,omitempty"`
    } `json:"invoince_confirm,omitempty"`
    

}
*/

// Invoinceconfirmlist 
type Invoinceconfirmlist struct {

    // 发票确认信息
    InvoinceConfirm   *Invoinceconfirm `json:"invoince_confirm,omitempty"`

}
