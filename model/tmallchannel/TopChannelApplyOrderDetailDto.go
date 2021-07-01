package tmallchannel

// TopChannelApplyOrderDetailDto 
type TopChannelApplyOrderDetailDto struct {
    // 申请单的宝贝详情
    ApplyOrderRelateItemList   []TopChannelApplyOrderRelateItemDto `json:"apply_order_relate_item_list,omitempty" xml:"apply_order_relate_item_list>top_channel_apply_order_relate_item_dto,omitempty"`
    // 整单价格（包括邮费，单位分）
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 整单邮费
    TotalPostFee   int64 `json:"total_post_fee,omitempty" xml:"total_post_fee,omitempty"`
}
