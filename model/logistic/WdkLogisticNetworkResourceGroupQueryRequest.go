package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询网格仓-区块-自提点关系 APIRequest
wdk.logistic.network.resource.group.query

查询网格仓-区块-自提点关系
*/
type WdkLogisticNetworkResourceGroupQueryRequest struct {
    model.Params

    // 入参
    paramResourceGroupPageQueryRequest   *ResourceGroupPageQueryRequest 

}

func NewWdkLogisticNetworkResourceGroupQueryRequest() *WdkLogisticNetworkResourceGroupQueryRequest{
    return &WdkLogisticNetworkResourceGroupQueryRequest{
        Params: model.NewParams(),
    }
}

func (r WdkLogisticNetworkResourceGroupQueryRequest) GetApiMethodName() string {
    return "wdk.logistic.network.resource.group.query"
}

func (r WdkLogisticNetworkResourceGroupQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *WdkLogisticNetworkResourceGroupQueryRequest) SetParamResourceGroupPageQueryRequest(paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest) error {
    r.paramResourceGroupPageQueryRequest = paramResourceGroupPageQueryRequest
    r.Set("param_resource_group_page_query_request", paramResourceGroupPageQueryRequest)
    return nil
}

func (r WdkLogisticNetworkResourceGroupQueryRequest) GetParamResourceGroupPageQueryRequest() *ResourceGroupPageQueryRequest {
    return r.paramResourceGroupPageQueryRequest
}

