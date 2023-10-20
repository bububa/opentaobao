package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest 每日优鲜仓作业单回传接口 API请求
// alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback
//
// 家乐福仓作业单回传接口
type AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *MissfreshO2ocallbackOrder
}

// NewAlibabawdkfulfillmissfreshwarehouseworkordercallbackRequest 初始化AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest对象
func NewAlibabawdkfulfillmissfreshwarehouseworkordercallbackRequest() *AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest {
	return &AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackOrder is CallbackOrder Setter
// 作业单回传对象
func (r *AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest) SetCallbackOrder(_callbackOrder *MissfreshO2ocallbackOrder) error {
	r._callbackOrder = _callbackOrder
	r.Set("callback_order", _callbackOrder)
	return nil
}

// GetCallbackOrder CallbackOrder Getter
func (r AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest) GetCallbackOrder() *MissfreshO2ocallbackOrder {
	return r._callbackOrder
}
