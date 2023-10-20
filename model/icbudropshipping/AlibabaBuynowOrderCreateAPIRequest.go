package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBuynowOrderCreateAPIRequest 阿里巴巴买家buynow下单接口 API请求
// alibaba.buynow.order.create
//
// 阿里巴巴买家下单接口
type AlibabaBuynowOrderCreateAPIRequest struct {
	model.Params
	// Order creation parameter
	_paramOrderCreateRequest *OrderCreateRequest
}

// NewAlibabaBuynowOrderCreateRequest 初始化AlibabaBuynowOrderCreateAPIRequest对象
func NewAlibabaBuynowOrderCreateRequest() *AlibabaBuynowOrderCreateAPIRequest {
	return &AlibabaBuynowOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBuynowOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.buynow.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBuynowOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBuynowOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderCreateRequest is ParamOrderCreateRequest Setter
// Order creation parameter
func (r *AlibabaBuynowOrderCreateAPIRequest) SetParamOrderCreateRequest(_paramOrderCreateRequest *OrderCreateRequest) error {
	r._paramOrderCreateRequest = _paramOrderCreateRequest
	r.Set("param_order_create_request", _paramOrderCreateRequest)
	return nil
}

// GetParamOrderCreateRequest ParamOrderCreateRequest Getter
func (r AlibabaBuynowOrderCreateAPIRequest) GetParamOrderCreateRequest() *OrderCreateRequest {
	return r._paramOrderCreateRequest
}
