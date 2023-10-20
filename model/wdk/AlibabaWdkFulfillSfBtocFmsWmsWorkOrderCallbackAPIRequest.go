package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest 顺丰仓作业回传 API请求
// alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback
//
// 顺丰仓作业单回传接口
type AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *SfB2cfmsCallbackOrder
}

// NewAlibabawdkfulfillsfbtocfmswmsworkordercallbackRequest 初始化AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest对象
func NewAlibabawdkfulfillsfbtocfmswmsworkordercallbackRequest() *AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest {
	return &AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest) SetCallbackOrder(_callbackOrder *SfB2cfmsCallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabawdkfulfillsfbtocfmswmsworkordercallbackAPIRequest) GetCallbackOrder() *SfB2cfmsCallbackOrder {
	return r._callbackOrder
}
