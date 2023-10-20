package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityUpdateAPIRequest 修改无条件单品优惠活动 API请求
// taobao.promotionmisc.item.activity.update
//
// 修改无条件单品优惠活动。&lt;br/&gt;1、该接口只修改活动基本信息和打折信息，如需要增加、删除参与该活动的商品，请调用taobao.promotionmisc.activity.range.add和taobao.promotionmisc.activity.range.remove接口。 &lt;br/&gt;2、使用该接口时需要同时把未做修改的字段值也传入。 &lt;br/&gt;&lt;br/&gt;3、该接口受店铺最低折扣限制，如优惠不生效，请让卖家检查该优惠是否低于店铺的最低折扣设置。
type TaobaoPromotionmiscItemActivityUpdateAPIRequest struct {
	model.Params
	// 活动名称。
	_name string
	// 活动开始时间。
	_startTime string
	// 活动结束时间。
	_endTime string
	// 用户标签。当is_user_tag为true时，该值才有意义。
	_userTag string
	// 活动id。
	_activityId int64
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

// NewTaobaoPromotionmiscItemActivityUpdateRequest 初始化TaobaoPromotionmiscItemActivityUpdateAPIRequest对象
func NewTaobaoPromotionmiscItemActivityUpdateRequest() *TaobaoPromotionmiscItemActivityUpdateAPIRequest {
	return &TaobaoPromotionmiscItemActivityUpdateAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) Reset() {
	r._name = ""
	r._startTime = ""
	r._endTime = ""
	r._userTag = ""
	r._activityId = 0
	r._participateRange = 0
	r._decreaseAmount = 0
	r._discountRate = 0
	r._isUserTag = false
	r._isDecreaseMoney = false
	r._isDiscount = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.promotionmisc.item.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 活动名称。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetName() string {
	return r._name
}

// SetStartTime is StartTime Setter
// 活动开始时间。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 活动结束时间。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetUserTag is UserTag Setter
// 用户标签。当is_user_tag为true时，该值才有意义。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetUserTag(_userTag string) error {
	r._userTag = _userTag
	r.Set("user_tag", _userTag)
	return nil
}

// GetUserTag UserTag Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetUserTag() string {
	return r._userTag
}

// SetActivityId is ActivityId Setter
// 活动id。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetParticipateRange is ParticipateRange Setter
// 活动范围：0表示全部参与； 1表示部分商品参与。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetParticipateRange(_participateRange int64) error {
	r._participateRange = _participateRange
	r.Set("participate_range", _participateRange)
	return nil
}

// GetParticipateRange ParticipateRange Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetParticipateRange() int64 {
	return r._participateRange
}

// SetDecreaseAmount is DecreaseAmount Setter
// 减多少钱。当is_decrease_money为true时，该值才有意义。注意：该值单位为分，即100表示1元。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetDecreaseAmount(_decreaseAmount int64) error {
	r._decreaseAmount = _decreaseAmount
	r.Set("decrease_amount", _decreaseAmount)
	return nil
}

// GetDecreaseAmount DecreaseAmount Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetDecreaseAmount() int64 {
	return r._decreaseAmount
}

// SetDiscountRate is DiscountRate Setter
// 折扣值。当is_discount为true时，该值才有意义。注意：800表示8折。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetDiscountRate(_discountRate int64) error {
	r._discountRate = _discountRate
	r.Set("discount_rate", _discountRate)
	return nil
}

// GetDiscountRate DiscountRate Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetDiscountRate() int64 {
	return r._discountRate
}

// SetIsUserTag is IsUserTag Setter
// 是否指定用户标签。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetIsUserTag(_isUserTag bool) error {
	r._isUserTag = _isUserTag
	r.Set("is_user_tag", _isUserTag)
	return nil
}

// GetIsUserTag IsUserTag Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetIsUserTag() bool {
	return r._isUserTag
}

// SetIsDecreaseMoney is IsDecreaseMoney Setter
// 是否有减钱行为。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetIsDecreaseMoney(_isDecreaseMoney bool) error {
	r._isDecreaseMoney = _isDecreaseMoney
	r.Set("is_decrease_money", _isDecreaseMoney)
	return nil
}

// GetIsDecreaseMoney IsDecreaseMoney Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetIsDecreaseMoney() bool {
	return r._isDecreaseMoney
}

// SetIsDiscount is IsDiscount Setter
// 是否有打折行为。
func (r *TaobaoPromotionmiscItemActivityUpdateAPIRequest) SetIsDiscount(_isDiscount bool) error {
	r._isDiscount = _isDiscount
	r.Set("is_discount", _isDiscount)
	return nil
}

// GetIsDiscount IsDiscount Getter
func (r TaobaoPromotionmiscItemActivityUpdateAPIRequest) GetIsDiscount() bool {
	return r._isDiscount
}

var poolTaobaoPromotionmiscItemActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionmiscItemActivityUpdateRequest()
	},
}

// GetTaobaoPromotionmiscItemActivityUpdateRequest 从 sync.Pool 获取 TaobaoPromotionmiscItemActivityUpdateAPIRequest
func GetTaobaoPromotionmiscItemActivityUpdateAPIRequest() *TaobaoPromotionmiscItemActivityUpdateAPIRequest {
	return poolTaobaoPromotionmiscItemActivityUpdateAPIRequest.Get().(*TaobaoPromotionmiscItemActivityUpdateAPIRequest)
}

// ReleaseTaobaoPromotionmiscItemActivityUpdateAPIRequest 将 TaobaoPromotionmiscItemActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionmiscItemActivityUpdateAPIRequest(v *TaobaoPromotionmiscItemActivityUpdateAPIRequest) {
	v.Reset()
	poolTaobaoPromotionmiscItemActivityUpdateAPIRequest.Put(v)
}
