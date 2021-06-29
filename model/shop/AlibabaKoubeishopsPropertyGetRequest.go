package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑店铺列表推荐 API请求
alibaba.koubeishops.property.get

推荐用户附近的美食门店
*/
type AlibabaKoubeishopsPropertyGetRequest struct {
    model.Params
    // 入参
    _paramOpenApiSearchRequest   *OpenApiSearchRequest
}

// 初始化AlibabaKoubeishopsPropertyGetRequest对象
func NewAlibabaKoubeishopsPropertyGetRequest() *AlibabaKoubeishopsPropertyGetRequest{
    return &AlibabaKoubeishopsPropertyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaKoubeishopsPropertyGetRequest) GetApiMethodName() string {
    return "alibaba.koubeishops.property.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaKoubeishopsPropertyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOpenApiSearchRequest Setter
// 入参
func (r *AlibabaKoubeishopsPropertyGetRequest) SetParamOpenApiSearchRequest(_paramOpenApiSearchRequest *OpenApiSearchRequest) error {
    r._paramOpenApiSearchRequest = _paramOpenApiSearchRequest
    r.Set("param_open_api_search_request", _paramOpenApiSearchRequest)
    return nil
}

// ParamOpenApiSearchRequest Getter
func (r AlibabaKoubeishopsPropertyGetRequest) GetParamOpenApiSearchRequest() *OpenApiSearchRequest {
    return r._paramOpenApiSearchRequest
}
