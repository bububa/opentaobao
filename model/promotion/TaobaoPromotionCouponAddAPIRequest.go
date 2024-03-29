package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionCouponAddAPIRequest 创建店铺优惠券接口 API请求
// taobao.promotion.coupon.add
//
// 创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
type TaobaoPromotionCouponAddAPIRequest struct {
	model.Params
	// 优惠券的生效时间
	_startTime string
	// 优惠券的截止日期
	_endTime string
	// 优惠券的面额，必须是3，5，10，20，50，100
	_denominations int64
	// 订单满多少元才能用这个优惠券，500就是满500元才能使用
	_condition int64
}

// NewTaobaoPromotionCouponAddRequest 初始化TaobaoPromotionCouponAddAPIRequest对象
func NewTaobaoPromotionCouponAddRequest() *TaobaoPromotionCouponAddAPIRequest {
	return &TaobaoPromotionCouponAddAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoPromotionCouponAddAPIRequest) Reset() {
	r._startTime = ""
	r._endTime = ""
	r._denominations = 0
	r._condition = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPromotionCouponAddAPIRequest) GetApiMethodName() string {
	return "taobao.promotion.coupon.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPromotionCouponAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPromotionCouponAddAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoPromotionCouponAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoPromotionCouponAddRequest()
	},
}

// GetTaobaoPromotionCouponAddRequest 从 sync.Pool 获取 TaobaoPromotionCouponAddAPIRequest
func GetTaobaoPromotionCouponAddAPIRequest() *TaobaoPromotionCouponAddAPIRequest {
	return poolTaobaoPromotionCouponAddAPIRequest.Get().(*TaobaoPromotionCouponAddAPIRequest)
}

// ReleaseTaobaoPromotionCouponAddAPIRequest 将 TaobaoPromotionCouponAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoPromotionCouponAddAPIRequest(v *TaobaoPromotionCouponAddAPIRequest) {
	v.Reset()
	poolTaobaoPromotionCouponAddAPIRequest.Put(v)
}
