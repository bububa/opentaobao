package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababuynowordercreateAPIRequest 阿里巴巴买家buynow下单接口 API请求
// alibaba.buynow.order.create
//
// 阿里巴巴买家下单接口
type AlibababuynowordercreateAPIRequest struct {
	model.Params
	// Order creation parameter
	_paramOrderCreateRequest *OrderCreateRequest
}

// NewAlibababuynowordercreateRequest 初始化AlibababuynowordercreateAPIRequest对象
func NewAlibababuynowordercreateRequest() *AlibababuynowordercreateAPIRequest {
	return &AlibababuynowordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababuynowordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.buynow.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababuynowordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababuynowordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderCreateRequest is ParamOrderCreateRequest Setter
// Order creation parameter
func (r *AlibababuynowordercreateAPIRequest) SetParamOrderCreateRequest(_paramOrderCreateRequest *OrderCreateRequest) error {
	r._paramOrderCreateRequest = _paramOrderCreateRequest
	r.Set("param_order_create_request", _paramOrderCreateRequest)
	return nil
}

// GetParamOrderCreateRequest ParamOrderCreateRequest Getter
func (r AlibababuynowordercreateAPIRequest) GetParamOrderCreateRequest() *OrderCreateRequest {
	return r._paramOrderCreateRequest
}
