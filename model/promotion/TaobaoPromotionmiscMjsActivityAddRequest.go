package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建满就送活动 APIRequest
taobao.promotionmisc.mjs.activity.add

创建满就送活动。<br/>1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。 2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。 3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
type TaobaoPromotionmiscMjsActivityAddRequest struct {
    model.Params

    // 活动名称。
    name   string 

    // 活动类型： 1表示商品级别的活动；2表示店铺级别的活动
    type   int64 

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

func NewTaobaoPromotionmiscMjsActivityAddRequest() *TaobaoPromotionmiscMjsActivityAddRequest{
    return &TaobaoPromotionmiscMjsActivityAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.mjs.activity.add"
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetType() int64 {
    return r.type
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetParticipateRange(participateRange int64) error {
    r.participateRange = participateRange
    r.Set("participate_range", participateRange)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetParticipateRange() int64 {
    return r.participateRange
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsAmountOver(isAmountOver bool) error {
    r.isAmountOver = isAmountOver
    r.Set("is_amount_over", isAmountOver)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsAmountOver() bool {
    return r.isAmountOver
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetTotalPrice(totalPrice int64) error {
    r.totalPrice = totalPrice
    r.Set("total_price", totalPrice)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetTotalPrice() int64 {
    return r.totalPrice
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsAmountMultiple(isAmountMultiple bool) error {
    r.isAmountMultiple = isAmountMultiple
    r.Set("is_amount_multiple", isAmountMultiple)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsAmountMultiple() bool {
    return r.isAmountMultiple
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsItemCountOver(isItemCountOver bool) error {
    r.isItemCountOver = isItemCountOver
    r.Set("is_item_count_over", isItemCountOver)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsItemCountOver() bool {
    return r.isItemCountOver
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetItemCount(itemCount int64) error {
    r.itemCount = itemCount
    r.Set("item_count", itemCount)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetItemCount() int64 {
    return r.itemCount
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsItemMultiple(isItemMultiple bool) error {
    r.isItemMultiple = isItemMultiple
    r.Set("is_item_multiple", isItemMultiple)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsItemMultiple() bool {
    return r.isItemMultiple
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsShopMember(isShopMember bool) error {
    r.isShopMember = isShopMember
    r.Set("is_shop_member", isShopMember)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsShopMember() bool {
    return r.isShopMember
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetShopMemberLevel(shopMemberLevel int64) error {
    r.shopMemberLevel = shopMemberLevel
    r.Set("shop_member_level", shopMemberLevel)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetShopMemberLevel() int64 {
    return r.shopMemberLevel
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsUserTag() bool {
    return r.isUserTag
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetUserTag() string {
    return r.userTag
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsDecreaseMoney(isDecreaseMoney bool) error {
    r.isDecreaseMoney = isDecreaseMoney
    r.Set("is_decrease_money", isDecreaseMoney)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsDecreaseMoney() bool {
    return r.isDecreaseMoney
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetDecreaseAmount(decreaseAmount int64) error {
    r.decreaseAmount = decreaseAmount
    r.Set("decrease_amount", decreaseAmount)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetDecreaseAmount() int64 {
    return r.decreaseAmount
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsDiscount(isDiscount bool) error {
    r.isDiscount = isDiscount
    r.Set("is_discount", isDiscount)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsDiscount() bool {
    return r.isDiscount
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetDiscountRate(discountRate int64) error {
    r.discountRate = discountRate
    r.Set("discount_rate", discountRate)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetDiscountRate() int64 {
    return r.discountRate
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsSendGift(isSendGift bool) error {
    r.isSendGift = isSendGift
    r.Set("is_send_gift", isSendGift)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsSendGift() bool {
    return r.isSendGift
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetGiftName(giftName string) error {
    r.giftName = giftName
    r.Set("gift_name", giftName)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetGiftName() string {
    return r.giftName
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetGiftId(giftId int64) error {
    r.giftId = giftId
    r.Set("gift_id", giftId)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetGiftId() int64 {
    return r.giftId
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetGiftUrl(giftUrl string) error {
    r.giftUrl = giftUrl
    r.Set("gift_url", giftUrl)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetGiftUrl() string {
    return r.giftUrl
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetIsFreePost(isFreePost bool) error {
    r.isFreePost = isFreePost
    r.Set("is_free_post", isFreePost)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetIsFreePost() bool {
    return r.isFreePost
}

func (r *TaobaoPromotionmiscMjsActivityAddRequest) SetExcludeArea(excludeArea string) error {
    r.excludeArea = excludeArea
    r.Set("exclude_area", excludeArea)
    return nil
}

func (r TaobaoPromotionmiscMjsActivityAddRequest) GetExcludeArea() string {
    return r.excludeArea
}

