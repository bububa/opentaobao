package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑店铺列表推荐 APIRequest
alibaba.koubeishops.property.get

推荐用户附近的美食门店
*/
type AlibabaKoubeishopsPropertyGetRequest struct {
    model.Params

    // 入参
    paramOpenApiSearchRequest   *OpenApiSearchRequest 

}

func NewAlibabaKoubeishopsPropertyGetRequest() *AlibabaKoubeishopsPropertyGetRequest{
    return &AlibabaKoubeishopsPropertyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaKoubeishopsPropertyGetRequest) GetApiMethodName() string {
    return "alibaba.koubeishops.property.get"
}

func (r AlibabaKoubeishopsPropertyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaKoubeishopsPropertyGetRequest) SetParamOpenApiSearchRequest(paramOpenApiSearchRequest *OpenApiSearchRequest) error {
    r.paramOpenApiSearchRequest = paramOpenApiSearchRequest
    r.Set("param_open_api_search_request", paramOpenApiSearchRequest)
    return nil
}

func (r AlibabaKoubeishopsPropertyGetRequest) GetParamOpenApiSearchRequest() *OpenApiSearchRequest {
    return r.paramOpenApiSearchRequest
}

