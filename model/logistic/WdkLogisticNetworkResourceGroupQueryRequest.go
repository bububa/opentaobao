package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询网格仓-区块-自提点关系 API请求
wdk.logistic.network.resource.group.query

查询网格仓-区块-自提点关系
*/
type WdkLogisticNetworkResourceGroupQueryRequest struct {
    model.Params
    // 入参
    _paramResourceGroupPageQueryRequest   *ResourceGroupPageQueryRequest
}

// 初始化WdkLogisticNetworkResourceGroupQueryRequest对象
func NewWdkLogisticNetworkResourceGroupQueryRequest() *WdkLogisticNetworkResourceGroupQueryRequest{
    return &WdkLogisticNetworkResourceGroupQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkResourceGroupQueryRequest) GetApiMethodName() string {
    return "wdk.logistic.network.resource.group.query"
}

// IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkResourceGroupQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamResourceGroupPageQueryRequest Setter
// 入参
func (r *WdkLogisticNetworkResourceGroupQueryRequest) SetParamResourceGroupPageQueryRequest(_paramResourceGroupPageQueryRequest *ResourceGroupPageQueryRequest) error {
    r._paramResourceGroupPageQueryRequest = _paramResourceGroupPageQueryRequest
    r.Set("param_resource_group_page_query_request", _paramResourceGroupPageQueryRequest)
    return nil
}

// ParamResourceGroupPageQueryRequest Getter
func (r WdkLogisticNetworkResourceGroupQueryRequest) GetParamResourceGroupPageQueryRequest() *ResourceGroupPageQueryRequest {
    return r._paramResourceGroupPageQueryRequest
}
