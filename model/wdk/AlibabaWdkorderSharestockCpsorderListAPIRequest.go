package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockcpsorderlistAPIRequest cps正向分销订单批量回流 API请求
// alibaba.wdkorder.sharestock.cpsorder.list
//
// cps正向分销订单批量回流
type AlibabawdkordersharestockcpsorderlistAPIRequest struct {
	model.Params
	// 入参
	_cpsOrderRequest *CpsOrderRequest
}

// NewAlibabawdkordersharestockcpsorderlistRequest 初始化AlibabawdkordersharestockcpsorderlistAPIRequest对象
func NewAlibabawdkordersharestockcpsorderlistRequest() *AlibabawdkordersharestockcpsorderlistAPIRequest {
	return &AlibabawdkordersharestockcpsorderlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersharestockcpsorderlistAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.cpsorder.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersharestockcpsorderlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersharestockcpsorderlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpsOrderRequest is CpsOrderRequest Setter
// 入参
func (r *AlibabawdkordersharestockcpsorderlistAPIRequest) SetCpsOrderRequest(_cpsOrderRequest *CpsOrderRequest) error {
	r._cpsOrderRequest = _cpsOrderRequest
	r.Set("cps_order_request", _cpsOrderRequest)
	return nil
}

// GetCpsOrderRequest CpsOrderRequest Getter
func (r AlibabawdkordersharestockcpsorderlistAPIRequest) GetCpsOrderRequest() *CpsOrderRequest {
	return r._cpsOrderRequest
}
