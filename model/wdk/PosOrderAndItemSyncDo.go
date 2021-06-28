package wdk

// PosOrderAndItemSyncDo 
/* model for simplify = false
type PosOrderAndItemSyncDo struct {

    // 商品信息
    
    ItemInfos  struct {
        ItemInfo  []ItemInfo `json:"item_info,omitempty"`
    } `json:"item_infos,omitempty"`
    

    // 订单流水信息
    
    OrderInfo  *struct {
        OrderInfoDo  *OrderInfoDo `json:"order_info_do,omitempty"`
    } `json:"order_info,omitempty"`
    

}
*/

// PosOrderAndItemSyncDo 
type PosOrderAndItemSyncDo struct {

    // 商品信息
    ItemInfos   []ItemInfo `json:"item_infos,omitempty"`

    // 订单流水信息
    OrderInfo   *OrderInfoDo `json:"order_info,omitempty"`

}
