package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopapplysearchAPIRequest 【商旅】搜索审批单列表 API请求
// alitrip.btrip.corpop.apply.search
//
// 【商旅】搜索审批单列表
type AlitripbtripcorpopapplysearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenIsvSearchRq
}

// NewAlitripbtripcorpopapplysearchRequest 初始化AlitripbtripcorpopapplysearchAPIRequest对象
func NewAlitripbtripcorpopapplysearchRequest() *AlitripbtripcorpopapplysearchAPIRequest {
	return &AlitripbtripcorpopapplysearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopapplysearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.apply.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopapplysearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopapplysearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcorpopapplysearchAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopapplysearchAPIRequest) GetRq() *OpenIsvSearchRq {
	return r._rq
}
