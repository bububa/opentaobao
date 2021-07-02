package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponAddAPIRequest 创建店铺优惠券接口 API请求
// taobao.promotion.coupon.add
//
// 创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
type TaobaoPromotionCouponAddAPIRequest struct {
	model.Params
	// 优惠券的面额，必须是3，5，10，20，50，100
	_denominations int64
	// 优惠券的截止日期
	_endTime string
	// 订单满多少元才能用这个优惠券，500就是满500元才能使用
	_condition int64
	// 优惠券的生效时间
	_startTime string
}

// NewTaobaoPromotionCouponAddRequest 初始化TaobaoPromotionCouponAddAPIRequest对象
func NewTaobaoPromotionCouponAddRequest() *TaobaoPromotionCouponAddAPIRequest {
	return &TaobaoPromotionCouponAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponAddAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDenominations is Denominations Setter
// 优惠券的面额，必须是3，5，10，20，50，100
func (r *TaobaoPromotionCouponAddAPIRequest) SetDenominations(_denominations int64) error {
	r._denominations = _denominations
	r.Set("denominations", _denominations)
	return nil
}

// GetDenominations Denominations Getter
func (r TaobaoPromotionCouponAddAPIRequest) GetDenominations() int64 {
	return r._denominations
}

// SetEndTime is EndTime Setter
// 优惠券的截止日期
func (r *TaobaoPromotionCouponAddAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoPromotionCouponAddAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCondition is Condition Setter
// 订单满多少元才能用这个优惠券，500就是满500元才能使用
func (r *TaobaoPromotionCouponAddAPIRequest) SetCondition(_condition int64) error {
	r._condition = _condition
	r.Set("condition", _condition)
	return nil
}

// GetCondition Condition Getter
func (r TaobaoPromotionCouponAddAPIRequest) GetCondition() int64 {
	return r._condition
}

// SetStartTime is StartTime Setter
// 优惠券的生效时间
func (r *TaobaoPromotionCouponAddAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoPromotionCouponAddAPIRequest) GetStartTime() string {
	return r._startTime
}
