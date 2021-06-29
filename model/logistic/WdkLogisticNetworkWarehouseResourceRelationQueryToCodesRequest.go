package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
按网格仓查中心仓（带缓存） API请求
wdk.logistic.network.warehouse.resource.relation.query.to.codes

盒马集市，网格仓查询中心仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest struct {
    model.Params
    // 入参
    paramYxWarehouseResourceRelationQueryRequest   *YxWarehouseResourceRelationQueryRequest
}

// 初始化WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest对象
func NewWdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest{
    return &WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest) GetApiMethodName() string {
    return "wdk.logistic.network.warehouse.resource.relation.query.to.codes"
}

// IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamYxWarehouseResourceRelationQueryRequest Setter
// 入参
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest) SetParamYxWarehouseResourceRelationQueryRequest(paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest) error {
    r.paramYxWarehouseResourceRelationQueryRequest = paramYxWarehouseResourceRelationQueryRequest
    r.Set("param_yx_warehouse_resource_relation_query_request", paramYxWarehouseResourceRelationQueryRequest)
    return nil
}

// ParamYxWarehouseResourceRelationQueryRequest Getter
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest) GetParamYxWarehouseResourceRelationQueryRequest() *YxWarehouseResourceRelationQueryRequest {
    return r.paramYxWarehouseResourceRelationQueryRequest
}
