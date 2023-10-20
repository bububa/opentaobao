package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopexceedapplysyncAPIRequest 第三方超标审批结果回传 API请求
// alitrip.btrip.corpop.exceedapply.sync
//
// 第三方审批单推送到企业后，企业审批结束，将审批结果回传给阿里商旅
type AlitripbtripcorpopexceedapplysyncAPIRequest struct {
	model.Params
	// 入参
	_rq *BtripExceedApplyRq
}

// NewAlitripbtripcorpopexceedapplysyncRequest 初始化AlitripbtripcorpopexceedapplysyncAPIRequest对象
func NewAlitripbtripcorpopexceedapplysyncRequest() *AlitripbtripcorpopexceedapplysyncAPIRequest {
	return &AlitripbtripcorpopexceedapplysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopexceedapplysyncAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.exceedapply.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopexceedapplysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopexceedapplysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripcorpopexceedapplysyncAPIRequest) SetRq(_rq *BtripExceedApplyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopexceedapplysyncAPIRequest) GetRq() *BtripExceedApplyRq {
	return r._rq
}
