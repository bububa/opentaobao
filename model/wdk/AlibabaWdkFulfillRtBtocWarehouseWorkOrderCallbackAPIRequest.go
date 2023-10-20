package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest 大润发B2C仓作业单回传接口 API请求
// alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback
//
// 大润发B2C仓作业单回传接口
type AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DrfB2ccallbackOrder
}

// NewAlibabawdkfulfillrtbtocwarehouseworkordercallbackRequest 初始化AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest对象
func NewAlibabawdkfulfillrtbtocwarehouseworkordercallbackRequest() *AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest {
	return &AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest) SetCallbackOrder(_callbackOrder *DrfB2ccallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest) GetCallbackOrder() *DrfB2ccallbackOrder {
	return r._callbackOrder
}
