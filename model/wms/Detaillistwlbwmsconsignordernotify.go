package wms

// Detaillistwlbwmsconsignordernotify 
/* model for simplify = false
type Detaillistwlbwmsconsignordernotify struct {

    // 发票信息
    
    ItemDetail  *struct {
        Itemdetailwlbwmsconsignordernotify  *Itemdetailwlbwmsconsignordernotify `json:"itemdetailwlbwmsconsignordernotify,omitempty"`
    } `json:"item_detail,omitempty"`
    

}
*/

// Detaillistwlbwmsconsignordernotify 
type Detaillistwlbwmsconsignordernotify struct {

    // 发票信息
    ItemDetail   *Itemdetailwlbwmsconsignordernotify `json:"item_detail,omitempty"`

}
