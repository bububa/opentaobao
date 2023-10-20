package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomobilepromotioncouponapplyAPIRequest 优惠券领取(手淘专用) API请求
// taobao.mobile.promotion.coupon.apply
//
// 优惠券领取
type TaobaomobilepromotioncouponapplyAPIRequest struct {
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

// NewTaobaomobilepromotioncouponapplyRequest 初始化TaobaomobilepromotioncouponapplyAPIRequest对象
func NewTaobaomobilepromotioncouponapplyRequest() *TaobaomobilepromotioncouponapplyAPIRequest {
	return &TaobaomobilepromotioncouponapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetApiMethodName() string {
	return "taobao.mobile.promotion.coupon.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 请求唯一id，问题排查
func (r *TaobaomobilepromotioncouponapplyAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetFeedId is FeedId Setter
// 广播id
func (r *TaobaomobilepromotioncouponapplyAPIRequest) SetFeedId(_feedId string) error {
	r._feedId = _feedId
	r.Set("feed_id", _feedId)
	return nil
}

// GetFeedId FeedId Getter
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetFeedId() string {
	return r._feedId
}

// SetBizId is BizId Setter
// 三方活动id
func (r *TaobaomobilepromotioncouponapplyAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetBizId() string {
	return r._bizId
}

// SetSpreadId is SpreadId Setter
// 传播id
func (r *TaobaomobilepromotioncouponapplyAPIRequest) SetSpreadId(_spreadId int64) error {
	r._spreadId = _spreadId
	r.Set("spread_id", _spreadId)
	return nil
}

// GetSpreadId SpreadId Getter
func (r TaobaomobilepromotioncouponapplyAPIRequest) GetSpreadId() int64 {
	return r._spreadId
}
