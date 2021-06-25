package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建无条件单品优惠活动 APIRequest
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

func NewTaobaoPromotionmiscItemActivityAddRequest() *TaobaoPromotionmiscItemActivityAddRequest{
    return &TaobaoPromotionmiscItemActivityAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetApiMethodName() string {
    return "taobao.promotionmisc.item.activity.add"
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionmiscItemActivityAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetParticipateRange(participateRange int64) error {
    r.participateRange = participateRange
    r.Set("participate_range", participateRange)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetParticipateRange() int64 {
    return r.participateRange
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetIsUserTag(isUserTag bool) error {
    r.isUserTag = isUserTag
    r.Set("is_user_tag", isUserTag)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetIsUserTag() bool {
    return r.isUserTag
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetUserTag(userTag string) error {
    r.userTag = userTag
    r.Set("user_tag", userTag)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetUserTag() string {
    return r.userTag
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetIsDecreaseMoney(isDecreaseMoney bool) error {
    r.isDecreaseMoney = isDecreaseMoney
    r.Set("is_decrease_money", isDecreaseMoney)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetIsDecreaseMoney() bool {
    return r.isDecreaseMoney
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetDecreaseAmount(decreaseAmount int64) error {
    r.decreaseAmount = decreaseAmount
    r.Set("decrease_amount", decreaseAmount)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetDecreaseAmount() int64 {
    return r.decreaseAmount
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetIsDiscount(isDiscount bool) error {
    r.isDiscount = isDiscount
    r.Set("is_discount", isDiscount)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetIsDiscount() bool {
    return r.isDiscount
}

func (r *TaobaoPromotionmiscItemActivityAddRequest) SetDiscountRate(discountRate int64) error {
    r.discountRate = discountRate
    r.Set("discount_rate", discountRate)
    return nil
}

func (r TaobaoPromotionmiscItemActivityAddRequest) GetDiscountRate() int64 {
    return r.discountRate
}

