package promotion

// LimitDiscountDetail 
type LimitDiscountDetail struct {
    // 限时打折名称
    LimitDiscountName   string `json:"limit_discount_name,omitempty" xml:"limit_discount_name,omitempty"`
    // 限时打折开始时间。
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 限时打折结束时间。
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 商品的id(数字类型)
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 该商品限时折扣
    ItemDiscount   float64 `json:"item_discount,omitempty" xml:"item_discount,omitempty"`
    // 每人限购数量，1、2、5、10000(不限)。
    LimitNum   int64 `json:"limit_num,omitempty" xml:"limit_num,omitempty"`
}
