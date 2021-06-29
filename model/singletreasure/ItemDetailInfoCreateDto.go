package singletreasure

// ItemDetailInfoCreateDto 
type ItemDetailInfoCreateDto struct {
    // 商品买家限购数量，-1表示不限购
    LimitCheck   int64 `json:"limit_check,omitempty" xml:"limit_check,omitempty"`
    // 是否取整
    IsMathFloor   bool `json:"is_math_floor,omitempty" xml:"is_math_floor,omitempty"`
    // 猫客折上折，优惠力度，打折、减钱：单位分；打折，8折：800
    MkDiscount   int64 `json:"mk_discount,omitempty" xml:"mk_discount,omitempty"`
    // sku对象列表
    SkuVOs   []SkuDetailInfoCreateDto `json:"sku_v_os,omitempty" xml:"sku_v_os>sku_detail_info_create_dto,omitempty"`
    // 商品Id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 是否抹分
    IsDiscardFen   bool `json:"is_discard_fen,omitempty" xml:"is_discard_fen,omitempty"`
    // 优惠力度，打折、减钱：单位分；打折，8折：800
    Discount   int64 `json:"discount,omitempty" xml:"discount,omitempty"`
    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
