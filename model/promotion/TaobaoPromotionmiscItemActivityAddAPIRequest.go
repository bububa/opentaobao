package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionmiscitemactivityaddAPIRequest 创建无条件单品优惠活动 API请求
// taobao.promotionmisc.item.activity.add
//
// 创建无条件单品优惠活动。1、可以选择是全店参加或者部分商品参加：participate_range：0表示全部参与； 1表示部分商品参与。&lt;br/&gt;2、如果是部分商品参加，则需要通过taobao.promotionmisc.activity.range.add接口来指定需要参加的商品。&lt;br/&gt;3、该接口创建的优惠受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
type TaobaopromotionmiscitemactivityaddAPIRequest struct {
	model.Params
	// 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
	_name string
	// 活动开始时间。
	_startTime string
	// 活动结束时间。
	_endTime string
	// 用户标签。当is_user_tag为true时，该值才有意义。
	_userTag string
	// 活动范围：0表示全部参与； 1表示部分商品参与。
	_participateRange int64
	// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
	_decreaseAmount int64
	// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
	_discountRate int64
	// 是否指定用户标签。
	_isUserTag bool
	// 是否有减钱行为。
	_isDecreaseMoney bool
	// 是否有打折行为。
	_isDiscount bool
}

// NewTaobaopromotionmiscitemactivityaddRequest 初始化TaobaopromotionmiscitemactivityaddAPIRequest对象
func NewTaobaopromotionmiscitemactivityaddRequest() *TaobaopromotionmiscitemactivityaddAPIRequest {
	return &TaobaopromotionmiscitemactivityaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 活动名称，超过5个汉字时，商品详情中显示的优惠名称为：卖家优惠。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetName() string {
	return r._name
}

// SetStartTime is StartTime Setter
// 活动开始时间。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetUserTag is UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetUserTag(_userTag string) error {
	r._userTag = _userTag
	r.Set("user_tag", _userTag)
	return nil
}

// GetUserTag UserTag Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetUserTag() string {
	return r._userTag
}

// SetParticipateRange is ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetParticipateRange(_participateRange int64) error {
	r._participateRange = _participateRange
	r.Set("participate_range", _participateRange)
	return nil
}

// GetParticipateRange ParticipateRange Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetParticipateRange() int64 {
	return r._participateRange
}

// SetDecreaseAmount is DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetDecreaseAmount(_decreaseAmount int64) error {
	r._decreaseAmount = _decreaseAmount
	r.Set("decrease_amount", _decreaseAmount)
	return nil
}

// GetDecreaseAmount DecreaseAmount Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetDecreaseAmount() int64 {
	return r._decreaseAmount
}

// SetDiscountRate is DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetDiscountRate(_discountRate int64) error {
	r._discountRate = _discountRate
	r.Set("discount_rate", _discountRate)
	return nil
}

// GetDiscountRate DiscountRate Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetDiscountRate() int64 {
	return r._discountRate
}

// SetIsUserTag is IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetIsUserTag(_isUserTag bool) error {
	r._isUserTag = _isUserTag
	r.Set("is_user_tag", _isUserTag)
	return nil
}

// GetIsUserTag IsUserTag Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetIsUserTag() bool {
	return r._isUserTag
}

// SetIsDecreaseMoney is IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetIsDecreaseMoney(_isDecreaseMoney bool) error {
	r._isDecreaseMoney = _isDecreaseMoney
	r.Set("is_decrease_money", _isDecreaseMoney)
	return nil
}

// GetIsDecreaseMoney IsDecreaseMoney Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetIsDecreaseMoney() bool {
	return r._isDecreaseMoney
}

// SetIsDiscount is IsDiscount Setter
// 是否有打折行为。
func (r *TaobaopromotionmiscitemactivityaddAPIRequest) SetIsDiscount(_isDiscount bool) error {
	r._isDiscount = _isDiscount
	r.Set("is_discount", _isDiscount)
	return nil
}

// GetIsDiscount IsDiscount Getter
func (r TaobaopromotionmiscitemactivityaddAPIRequest) GetIsDiscount() bool {
	return r._isDiscount
}
