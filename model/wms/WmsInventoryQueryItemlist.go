package wms

// WmsInventoryQueryItemlist 
/* model for simplify = false
type WmsInventoryQueryItemlist struct {

    // 商品详情
    
    Item  *struct {
        WmsInventoryQueryItem  *WmsInventoryQueryItem `json:"wms_inventory_query_item,omitempty"`
    } `json:"item,omitempty"`
    

}
*/

// WmsInventoryQueryItemlist 
type WmsInventoryQueryItemlist struct {

    // 商品详情
    Item   *WmsInventoryQueryItem `json:"item,omitempty"`

}
