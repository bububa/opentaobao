package btrip

// BtripHotelValidateOrderRq 
type BtripHotelValidateOrderRq struct {

    // 购买人在分销商平台的用户昵称
    
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    

    // 购买人在分销商平台的唯一用户id
    
    BuyerUniqueKey   string `json:"buyer_unique_key,omitempty" xml:"buyer_unique_key,omitempty"`
    

    // 入住时间
    
    CheckIn   string `json:"check_in,omitempty" xml:"check_in,omitempty"`
    

    // 离店时间
    
    CheckOut   string `json:"check_out,omitempty" xml:"check_out,omitempty"`
    

    // 商旅酒店唯一商品标识
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 每间房成人数
    
    NumberOfAdultsPerRoom   int64 `json:"number_of_adults_per_room,omitempty" xml:"number_of_adults_per_room,omitempty"`
    

    // 预订房间数量
    
    NumberOfRooms   int64 `json:"number_of_rooms,omitempty" xml:"number_of_rooms,omitempty"`
    

    // 销计划id
    
    RatePlanId   int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
    

    // 详情报价的优惠金额
    
    SearchPromotionAmount   int64 `json:"search_promotion_amount,omitempty" xml:"search_promotion_amount,omitempty"`
    

    // 详情报价的房价
    
    SearchRoomPrice   int64 `json:"search_room_price,omitempty" xml:"search_room_price,omitempty"`
    

    // 分销子渠道，通常指代商旅的企业id
    
    SubChannel   string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
    

    // 供应商标识
    
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    

    // 总价
    
    TotalPrice   int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
    

}
