package drug

// SuitSubItemDto 
type SuitSubItemDto struct {

    // 单位
    Unit   string `json:"unit,omitempty"`

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 个数
    Quantity   int64 `json:"quantity,omitempty"`

    // 药标题
    Title   string `json:"title,omitempty"`

    // 商家自定义商品ID
    OutId   string `json:"out_id,omitempty"`

}
