package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest 北京小蜜蜂配作业回传 API请求
// alibaba.wdk.fulfill.dms.ebeecake.work.order.callback
//
// 北京小蜜蜂配作业回传。
type AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *EbeecakeO2ocallbackOrder
}

// NewAlibabawdkfulfilldmsebeecakeworkordercallbackRequest 初始化AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest对象
func NewAlibabawdkfulfilldmsebeecakeworkordercallbackRequest() *AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest {
	return &AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.dms.ebeecake.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest) SetCallbackOrder(_callbackOrder *EbeecakeO2ocallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest) GetCallbackOrder() *EbeecakeO2ocallbackOrder {
	return r._callbackOrder
}
