package promotion

// LimitDiscount 
type LimitDiscount struct {
    // 限时打折ID。
    LimitDiscountId   int64 `json:"limit_discount_id,omitempty" xml:"limit_discount_id,omitempty"`
    // 限时打折名称。
    LimitDiscountName   string `json:"limit_discount_name,omitempty" xml:"limit_discount_name,omitempty"`
    // 限时打折开始时间。
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 限时打折结束时间。
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 该限时打折宝贝数量。
    ItemNum   int64 `json:"item_num,omitempty" xml:"item_num,omitempty"`
}
