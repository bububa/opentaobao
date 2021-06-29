package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建无条件单品优惠活动 API请求
taobao.promotionmisc.item.activity.add

创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。<br/>2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。<br/>3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
*/
type TaobaoPromotionmiscItemActivityAddRequest struct {
    model.Params
    // 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
    name   string
    // 活动范围：0表示全部参与； 1表示部分商品参与。
    participateRange   int64
    // 活动开始时间。
    startTime   string
    // 活动结束时间。
    endTime   string
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
}

// 初始化TaobaoPromotionmiscItemActivityAddRequest对象
func NewTaobaoPromotionmiscItemActivityAddRequest() *TaobaoPromotionmiscItemActivityAddRequest{
    return &TaobaoPromotionmiscItemActivityAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetName() string {
    return r.name
}
// ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetParticipateRange(participateRange int64) error {
    r.participateRange = participateRange
    r.Set("participate_range", participateRange)
    return nil
}

// ParticipateRange Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetParticipateRange() int64 {
    return r.participateRange
}
// StartTime Setter
// 活动开始时间。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 活动结束时间。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetEndTime() string {
    return r.endTime
}
// IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

// IsUserTag Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetIsUserTag() bool {
    return r.isUserTag
}
// UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

// UserTag Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetUserTag() string {
    return r.userTag
}
// IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetIsDecreaseMoney(isDecreaseMoney bool) error {
    r.isDecreaseMoney = isDecreaseMoney
    r.Set("is_decrease_money", isDecreaseMoney)
    return nil
}

// IsDecreaseMoney Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetIsDecreaseMoney() bool {
    return r.isDecreaseMoney
}
// DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetDecreaseAmount(decreaseAmount int64) error {
    r.decreaseAmount = decreaseAmount
    r.Set("decrease_amount", decreaseAmount)
    return nil
}

// DecreaseAmount Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetDecreaseAmount() int64 {
    return r.decreaseAmount
}
// IsDiscount Setter
// 是否有打折行为。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetIsDiscount(isDiscount bool) error {
    r.isDiscount = isDiscount
    r.Set("is_discount", isDiscount)
    return nil
}

// IsDiscount Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetIsDiscount() bool {
    return r.isDiscount
}
// DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaoPromotionmiscItemActivityAddRequest) SetDiscountRate(discountRate int64) error {
    r.discountRate = discountRate
    r.Set("discount_rate", discountRate)
    return nil
}

// DiscountRate Getter
func (r TaobaoPromotionmiscItemActivityAddRequest) GetDiscountRate() int64 {
    return r.discountRate
}
