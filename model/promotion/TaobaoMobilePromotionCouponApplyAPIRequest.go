package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMobilePromotionCouponApplyAPIRequest 优惠券领取(手淘专用) API请求
// taobao.mobile.promotion.coupon.apply
//
// 优惠券领取
type TaobaoMobilePromotionCouponApplyAPIRequest struct {
	model.Params
	// 请求唯一id，问题排查
	_traceId string
	// 广播id
	_feedId string
	// 三方活动id
	_bizId string
	// 传播id
	_spreadId int64
}

// NewTaobaoMobilePromotionCouponApplyRequest 初始化TaobaoMobilePromotionCouponApplyAPIRequest对象
func NewTaobaoMobilePromotionCouponApplyRequest() *TaobaoMobilePromotionCouponApplyAPIRequest {
	return &TaobaoMobilePromotionCouponApplyAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) Reset() {
	r._traceId = ""
	r._feedId = ""
	r._bizId = ""
	r._spreadId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 请求唯一id，问题排查
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetFeedId is FeedId Setter
// 广播id
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetFeedId(_feedId string) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetFeedId() string {
	return r._feedId
}

// SetBizId is BizId Setter
// 三方活动id
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetBizId() string {
	return r._bizId
}

// SetSpreadId is SpreadId Setter
// 传播id
func (r *TaobaoMobilePromotionCouponApplyAPIRequest) SetSpreadId(_spreadId int64) error {
	r._spreadId = _spreadId
	r.Set("spread_id", _spreadId)
	return nil
}

// GetSpreadId SpreadId Getter
func (r TaobaoMobilePromotionCouponApplyAPIRequest) GetSpreadId() int64 {
	return r._spreadId
}

var poolTaobaoMobilePromotionCouponApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMobilePromotionCouponApplyRequest()
	},
}

// GetTaobaoMobilePromotionCouponApplyRequest 从 sync.Pool 获取 TaobaoMobilePromotionCouponApplyAPIRequest
func GetTaobaoMobilePromotionCouponApplyAPIRequest() *TaobaoMobilePromotionCouponApplyAPIRequest {
	return poolTaobaoMobilePromotionCouponApplyAPIRequest.Get().(*TaobaoMobilePromotionCouponApplyAPIRequest)
}

// ReleaseTaobaoMobilePromotionCouponApplyAPIRequest 将 TaobaoMobilePromotionCouponApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoMobilePromotionCouponApplyAPIRequest(v *TaobaoMobilePromotionCouponApplyAPIRequest) {
	v.Reset()
	poolTaobaoMobilePromotionCouponApplyAPIRequest.Put(v)
}
