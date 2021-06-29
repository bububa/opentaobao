package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改满就送活动 API请求
taobao.promotionmisc.mjs.activity.update

修改满就送活动。 <br/>1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 <br/>2、使用该接口时需要同时把未做修改的字段值也传入。 <br/>3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
type TaobaoPromotionmiscMjsActivityUpdateRequest struct {
    model.Params
    // 活动id。
    activityId   int64
    // 活动名称。
    name   string
    // 活动范围：0表示全部参与； 1表示部分商品参与。
    participateRange   int64
    // 活动开始时间。
    startTime   string
    // 活动结束时间。
    endTime   string
    // 是否有满元条件。
    isAmountOver   bool
    // 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
    totalPrice   int64
    // 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。默认为false。
    isAmountMultiple   bool
    // 是否有满件条件。
    isItemCountOver   bool
    // 满多少件。当is_item_count_over为true时，该值才有意义。
    itemCount   int64
    // 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。 默认为false。
    isItemMultiple   bool
    // 是否有店铺会员等级条件。
    isShopMember   bool
    // 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
    shopMemberLevel   int64
    // 是否指定用户标签。
    isUserTag   bool
    // 用户标签。当is_user_tag为true时，该值才有意义。
    userTag   string
    // 是否有减钱行为。
    isDecreaseMoney   bool
    // 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
    decreaseAmount   int64
    // 是否有打折行为。
    isDiscount   bool
    // 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
    discountRate   int64
    // 是否有送礼品行为。
    isSendGift   bool
    // 礼品名称。当is_send_gift为true时，该值才有意义。
    giftName   string
    // 礼品id，当is_send_gift为true时，该值才有意义。 1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
    giftId   int64
    // 商品详情的url，当is_send_gift为true时，该值才有效。
    giftUrl   string
    // 是否有免邮行为。
    isFreePost   bool
    // 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
    excludeArea   string
}

// 初始化TaobaoPromotionmiscMjsActivityUpdateRequest对象
func NewTaobaoPromotionmiscMjsActivityUpdateRequest() *TaobaoPromotionmiscMjsActivityUpdateRequest{
    return &TaobaoPromotionmiscMjsActivityUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetActivityId(activityId int64) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetActivityId() int64 {
    return r.activityId
}
// Name Setter
// 活动名称。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetName() string {
    return r.name
}
// ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetParticipateRange(participateRange int64) error {
    r.participateRange = participateRange
    r.Set("participate_range", participateRange)
    return nil
}

// ParticipateRange Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetParticipateRange() int64 {
    return r.participateRange
}
// StartTime Setter
// 活动开始时间。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 活动结束时间。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetEndTime() string {
    return r.endTime
}
// IsAmountOver Setter
// 是否有满元条件。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsAmountOver(isAmountOver bool) error {
    r.isAmountOver = isAmountOver
    r.Set("is_amount_over", isAmountOver)
    return nil
}

// IsAmountOver Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsAmountOver() bool {
    return r.isAmountOver
}
// TotalPrice Setter
// 满多少元。当is_amount_over为true时，该才字段有意义。注意：单位是分，即10000表示100元。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetTotalPrice(totalPrice int64) error {
    r.totalPrice = totalPrice
    r.Set("total_price", totalPrice)
    return nil
}

// TotalPrice Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetTotalPrice() int64 {
    return r.totalPrice
}
// IsAmountMultiple Setter
// 满元是否上不封顶。当is_amount_over为true时，该值才有意义。当该值为true时，表示满元上不封顶，例如满100元减10元，当满200时，则减20元。。。默认为false。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsAmountMultiple(isAmountMultiple bool) error {
    r.isAmountMultiple = isAmountMultiple
    r.Set("is_amount_multiple", isAmountMultiple)
    return nil
}

// IsAmountMultiple Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsAmountMultiple() bool {
    return r.isAmountMultiple
}
// IsItemCountOver Setter
// 是否有满件条件。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsItemCountOver(isItemCountOver bool) error {
    r.isItemCountOver = isItemCountOver
    r.Set("is_item_count_over", isItemCountOver)
    return nil
}

// IsItemCountOver Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsItemCountOver() bool {
    return r.isItemCountOver
}
// ItemCount Setter
// 满多少件。当is_item_count_over为true时，该值才有意义。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetItemCount(itemCount int64) error {
    r.itemCount = itemCount
    r.Set("item_count", itemCount)
    return nil
}

// ItemCount Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetItemCount() int64 {
    return r.itemCount
}
// IsItemMultiple Setter
// 满件是否上不封顶。当is_amount_multiple为true时，该值才有意义。当该值为true时，表示满件上不封顶，例如满10件减2元，当满20件时，则减4元。。。 默认为false。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsItemMultiple(isItemMultiple bool) error {
    r.isItemMultiple = isItemMultiple
    r.Set("is_item_multiple", isItemMultiple)
    return nil
}

// IsItemMultiple Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsItemMultiple() bool {
    return r.isItemMultiple
}
// IsShopMember Setter
// 是否有店铺会员等级条件。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsShopMember(isShopMember bool) error {
    r.isShopMember = isShopMember
    r.Set("is_shop_member", isShopMember)
    return nil
}

// IsShopMember Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsShopMember() bool {
    return r.isShopMember
}
// ShopMemberLevel Setter
// 店铺会员等级，当is_shop_member为true时，该值才有意义。0：店铺客户；1：普通客户；2：高级会员；3：VIP会员； 4：至尊VIP会员。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetShopMemberLevel(shopMemberLevel int64) error {
    r.shopMemberLevel = shopMemberLevel
    r.Set("shop_member_level", shopMemberLevel)
    return nil
}

// ShopMemberLevel Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetShopMemberLevel() int64 {
    return r.shopMemberLevel
}
// IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

// IsUserTag Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsUserTag() bool {
    return r.isUserTag
}
// UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

// UserTag Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetUserTag() string {
    return r.userTag
}
// IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsDecreaseMoney(isDecreaseMoney bool) error {
    r.isDecreaseMoney = isDecreaseMoney
    r.Set("is_decrease_money", isDecreaseMoney)
    return nil
}

// IsDecreaseMoney Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsDecreaseMoney() bool {
    return r.isDecreaseMoney
}
// DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetDecreaseAmount(decreaseAmount int64) error {
    r.decreaseAmount = decreaseAmount
    r.Set("decrease_amount", decreaseAmount)
    return nil
}

// DecreaseAmount Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetDecreaseAmount() int64 {
    return r.decreaseAmount
}
// IsDiscount Setter
// 是否有打折行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsDiscount(isDiscount bool) error {
    r.isDiscount = isDiscount
    r.Set("is_discount", isDiscount)
    return nil
}

// IsDiscount Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsDiscount() bool {
    return r.isDiscount
}
// DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetDiscountRate(discountRate int64) error {
    r.discountRate = discountRate
    r.Set("discount_rate", discountRate)
    return nil
}

// DiscountRate Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetDiscountRate() int64 {
    return r.discountRate
}
// IsSendGift Setter
// 是否有送礼品行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsSendGift(isSendGift bool) error {
    r.isSendGift = isSendGift
    r.Set("is_send_gift", isSendGift)
    return nil
}

// IsSendGift Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsSendGift() bool {
    return r.isSendGift
}
// GiftName Setter
// 礼品名称。当is_send_gift为true时，该值才有意义。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetGiftName(giftName string) error {
    r.giftName = giftName
    r.Set("gift_name", giftName)
    return nil
}

// GiftName Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetGiftName() string {
    return r.giftName
}
// GiftId Setter
// 礼品id，当is_send_gift为true时，该值才有意义。 1）只有填写真实的淘宝商品id时，才能生成物流单，并且在确定订单的页面上可以点击该商品名称跳转到商品详情页面。2）当礼物为实物商品时(有宝贝id),礼物必须为上架商品,不能为虚拟商品,不能为拍卖商品,不能有sku,不符合条件的,不做为礼物。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetGiftId(giftId int64) error {
    r.giftId = giftId
    r.Set("gift_id", giftId)
    return nil
}

// GiftId Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetGiftId() int64 {
    return r.giftId
}
// GiftUrl Setter
// 商品详情的url，当is_send_gift为true时，该值才有效。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetGiftUrl(giftUrl string) error {
    r.giftUrl = giftUrl
    r.Set("gift_url", giftUrl)
    return nil
}

// GiftUrl Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetGiftUrl() string {
    return r.giftUrl
}
// IsFreePost Setter
// 是否有免邮行为。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetIsFreePost(isFreePost bool) error {
    r.isFreePost = isFreePost
    r.Set("is_free_post", isFreePost)
    return nil
}

// IsFreePost Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetIsFreePost() bool {
    return r.isFreePost
}
// ExcludeArea Setter
// 免邮的排除地区，即，除指定地区外，其他地区包邮。当is_free_post为true时，该值才有意义。代码使用*链接，代码为行政区划代码。
func (r *TaobaoPromotionmiscMjsActivityUpdateRequest) SetExcludeArea(excludeArea string) error {
    r.excludeArea = excludeArea
    r.Set("exclude_area", excludeArea)
    return nil
}

// ExcludeArea Getter
func (r TaobaoPromotionmiscMjsActivityUpdateRequest) GetExcludeArea() string {
    return r.excludeArea
}
