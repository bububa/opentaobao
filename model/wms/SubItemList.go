package wms

// SubItemList 
/* model for simplify = false
type SubItemList struct {

    // 子货品
    
    SubItem  *struct {
        SubItem  *SubItem `json:"sub_item,omitempty"`
    } `json:"sub_item,omitempty"`
    

}
*/

// SubItemList 
type SubItemList struct {

    // 子货品
    SubItem   *SubItem `json:"sub_item,omitempty"`

}
