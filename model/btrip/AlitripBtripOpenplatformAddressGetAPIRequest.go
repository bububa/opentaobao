package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenplatformAddressGetAPIRequest 【商旅】开放平台对外页面跳转 API请求
// alitrip.btrip.openplatform.address.get
//
// 获取类目预定页跳转地址
type AlitripBtripOpenplatformAddressGetAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiJumpInfoRq
}

// NewAlitripBtripOpenplatformAddressGetRequest 初始化AlitripBtripOpenplatformAddressGetAPIRequest对象
func NewAlitripBtripOpenplatformAddressGetRequest() *AlitripBtripOpenplatformAddressGetAPIRequest {
	return &AlitripBtripOpenplatformAddressGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.openplatform.address.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripOpenplatformAddressGetAPIRequest) SetRq(_rq *OpenApiJumpInfoRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenplatformAddressGetAPIRequest) GetRq() *OpenApiJumpInfoRq {
	return r._rq
}
