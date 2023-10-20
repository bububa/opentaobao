package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopenplatformaddressgetAPIRequest 【商旅】开放平台对外页面跳转 API请求
// alitrip.btrip.openplatform.address.get
//
// 获取类目预定页跳转地址
type AlitripbtripopenplatformaddressgetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiJumpInfoRq
}

// NewAlitripbtripopenplatformaddressgetRequest 初始化AlitripbtripopenplatformaddressgetAPIRequest对象
func NewAlitripbtripopenplatformaddressgetRequest() *AlitripbtripopenplatformaddressgetAPIRequest {
	return &AlitripbtripopenplatformaddressgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopenplatformaddressgetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.openplatform.address.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopenplatformaddressgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopenplatformaddressgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripopenplatformaddressgetAPIRequest) SetRq(_rq *OpenApiJumpInfoRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopenplatformaddressgetAPIRequest) GetRq() *OpenApiJumpInfoRq {
	return r._rq
}
