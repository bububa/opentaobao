package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaKoubeishopsPropertyGetAPIRequest 口碑店铺列表推荐 API请求
// alibaba.koubeishops.property.get
//
// 推荐用户附近的美食门店
type AlibabaKoubeishopsPropertyGetAPIRequest struct {
	model.Params
	// 入参
	_paramOpenApiSearchRequest *OpenApiSearchRequest
}

// NewAlibabaKoubeishopsPropertyGetRequest 初始化AlibabaKoubeishopsPropertyGetAPIRequest对象
func NewAlibabaKoubeishopsPropertyGetRequest() *AlibabaKoubeishopsPropertyGetAPIRequest {
	return &AlibabaKoubeishopsPropertyGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetApiMethodName() string {
	return "alibaba.koubeishops.property.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamOpenApiSearchRequest Setter
// 入参
func (r *AlibabaKoubeishopsPropertyGetAPIRequest) SetParamOpenApiSearchRequest(_paramOpenApiSearchRequest *OpenApiSearchRequest) error {
	r._paramOpenApiSearchRequest = _paramOpenApiSearchRequest
	r.Set("param_open_api_search_request", _paramOpenApiSearchRequest)
	return nil
}

// Get ParamOpenApiSearchRequest Getter
func (r AlibabaKoubeishopsPropertyGetAPIRequest) GetParamOpenApiSearchRequest() *OpenApiSearchRequest {
	return r._paramOpenApiSearchRequest
}
