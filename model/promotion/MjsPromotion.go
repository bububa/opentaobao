package promotion

// MjsPromotion 
type MjsPromotion struct {

    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 活动名称。
    Name   string `json:"name,omitempty"`

    // 活动的有效条件、人群和行为描述。
    Description   string `json:"description,omitempty"`

    // 活动类型： 1表示商品级别的活动；2表示店铺级别的活动。
    Type   int64 `json:"type,omitempty"`

    // 活动范围：0表示全部参与； 1表示部分商品参与。
    ParticipateRange   int64 `json:"participate_range,omitempty"`

    // 活动开始时间。
    StartTime   string `json:"start_time,omitempty"`

    // 活动结束时间。
    EndTime   string `json:"end_time,omitempty"`

    // 是否有满元条件。
    IsAmountOver   bool `json:"is_amount_over,omitempty"`

    // 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
    TotalPrice   int64 `json:"total_price,omitempty"`

    // 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。
    IsAmountMultiple   bool `json:"is_amount_multiple,omitempty"`

    // 是否有满件条件。
    IsItemCountOver   bool `json:"is_item_count_over,omitempty"`

    // 满多少件。当is_item_count_over为true时，该值才有意义。
    ItemCount   int64 `json:"item_count,omitempty"`

    // 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。
    IsItemMultiple   bool `json:"is_item_multiple,omitempty"`

    // 是否有店铺会员等级条件。
    IsShopMember   bool `json:"is_shop_member,omitempty"`

    // 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
    ShopMemberLevel   int64 `json:"shop_member_level,omitempty"`

    // 是否指定用户标签。
    IsUserTag   bool `json:"is_user_tag,omitempty"`

    // 用户标签。当is_user_tag为true时，该值才有意义。
    UserTag   string `json:"user_tag,omitempty"`

    // 是否有减钱行为。
    IsDecreaseMoney   bool `json:"is_decrease_money,omitempty"`

    // 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
    DecreaseAmount   int64 `json:"decrease_amount,omitempty"`

    // 是否有打折行为。
    IsDiscount   bool `json:"is_discount,omitempty"`

    // 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
    DiscountRate   int64 `json:"discount_rate,omitempty"`

    // 是否有送礼品行为。
    IsSendGift   bool `json:"is_send_gift,omitempty"`

    // 礼品名称。当is_send_gift为true时，该值才有意义。
    GiftName   string `json:"gift_name,omitempty"`

    // 礼品id，当is_send_gift为true时，该值才有意义。<br/>1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
    GiftId   int64 `json:"gift_id,omitempty"`

    // 商品详情的url，当is_send_gift为true时，该值才有效。
    GiftUrl   string `json:"gift_url,omitempty"`

    // 是否有免邮行为。
    IsFreePost   bool `json:"is_free_post,omitempty"`

    // 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
    ExcludeArea   string `json:"exclude_area,omitempty"`

}
