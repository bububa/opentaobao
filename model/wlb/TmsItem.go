package wlb

// TmsItem 
type TmsItem struct {

    // 数量
    ItemQuantity   int64 `json:"item_quantity,omitempty"`

    // 前端商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 货品ID
    ScItemId   string `json:"sc_item_id,omitempty"`

}
