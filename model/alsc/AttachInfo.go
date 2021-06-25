package alsc

// AttachInfo 
type AttachInfo struct {

    // 附加商品实收金额
    ActualFee   int64 `json:"actual_fee,omitempty"`

    // 附加商品数量
    ItemCount   int64 `json:"item_count,omitempty"`

    // 附加商品id
    OutAttachItemId   string `json:"out_attach_item_id,omitempty"`

    // 附加商品名称
    OutAttachItemName   string `json:"out_attach_item_name,omitempty"`

    // 单价
    Price   int64 `json:"price,omitempty"`

    // 单位
    Unit   string `json:"unit,omitempty"`

    // 重量
    Weight   string `json:"weight,omitempty"`

    // 附加商品总金额
    TotalFee   int64 `json:"total_fee,omitempty"`

}
