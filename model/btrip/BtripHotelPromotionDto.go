package btrip

// BtripHotelPromotionDTO 
type BtripHotelPromotionDTO struct {
    // 详细的优惠信息列表
    PromotionDetailList   []BtripHotelPromotionDetailDTO `json:"promotion_detail_list,omitempty" xml:"promotion_detail_list>btrip_hotel_promotion_detail_dto,omitempty"`
    // 当前下单是否存在优惠
    PromotionExisted   bool `json:"promotion_existed,omitempty" xml:"promotion_existed,omitempty"`
    // 总优惠金额
    PromotionTotalPrice   int64 `json:"promotion_total_price,omitempty" xml:"promotion_total_price,omitempty"`
}
