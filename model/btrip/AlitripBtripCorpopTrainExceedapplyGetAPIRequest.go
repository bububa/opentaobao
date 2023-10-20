package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpoptrainexceedapplygetAPIRequest 商旅火车票第三方超标审批单搜索接口 API请求
// alitrip.btrip.corpop.train.exceedapply.get
//
// 商旅火车票第三方超标审批单搜索接口
type AlitripbtripcorpoptrainexceedapplygetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvSearchRq
}

// NewAlitripbtripcorpoptrainexceedapplygetRequest 初始化AlitripbtripcorpoptrainexceedapplygetAPIRequest对象
func NewAlitripbtripcorpoptrainexceedapplygetRequest() *AlitripbtripcorpoptrainexceedapplygetAPIRequest {
	return &AlitripbtripcorpoptrainexceedapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpoptrainexceedapplygetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.train.exceedapply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpoptrainexceedapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpoptrainexceedapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripcorpoptrainexceedapplygetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpoptrainexceedapplygetAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
