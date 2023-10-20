package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabakoubeishopspropertygetAPIRequest 口碑店铺列表推荐 API请求
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
type AlibabakoubeishopspropertygetAPIRequest struct {
	model.Params
	// 入参
	_paramOpenApiSearchRequest *OpenApiSearchRequest
}

// NewAlibabakoubeishopspropertygetRequest 初始化AlibabakoubeishopspropertygetAPIRequest对象
func NewAlibabakoubeishopspropertygetRequest() *AlibabakoubeishopspropertygetAPIRequest {
	return &AlibabakoubeishopspropertygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabakoubeishopspropertygetAPIRequest) GetApiMethodName() string {
	return "alibaba.koubeishops.property.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabakoubeishopspropertygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabakoubeishopspropertygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOpenApiSearchRequest is ParamOpenApiSearchRequest Setter
// 入参
func (r *AlibabakoubeishopspropertygetAPIRequest) SetParamOpenApiSearchRequest(_paramOpenApiSearchRequest *OpenApiSearchRequest) error {
	r._paramOpenApiSearchRequest = _paramOpenApiSearchRequest
	r.Set("param_open_api_search_request", _paramOpenApiSearchRequest)
	return nil
}

// GetParamOpenApiSearchRequest ParamOpenApiSearchRequest Getter
func (r AlibabakoubeishopspropertygetAPIRequest) GetParamOpenApiSearchRequest() *OpenApiSearchRequest {
	return r._paramOpenApiSearchRequest
}
