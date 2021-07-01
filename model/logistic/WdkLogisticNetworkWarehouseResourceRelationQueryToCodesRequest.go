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
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest struct {
    model.Params
    // 入参
    _paramYxWarehouseResourceRelationQueryRequest   *YxWarehouseResourceRelationQueryRequest
}

// 初始化WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest对象
func NewWdkLogisticNetworkWarehouseResourceRelationQueryToCodesRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest{
    return &WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetApiMethodName() string {
    return "wdk.logistic.network.warehouse.resource.relation.query.to.codes"
}

// IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamYxWarehouseResourceRelationQueryRequest Setter
// 入参
func (r *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) SetParamYxWarehouseResourceRelationQueryRequest(_paramYxWarehouseResourceRelationQueryRequest *YxWarehouseResourceRelationQueryRequest) error {
    r._paramYxWarehouseResourceRelationQueryRequest = _paramYxWarehouseResourceRelationQueryRequest
    r.Set("param_yx_warehouse_resource_relation_query_request", _paramYxWarehouseResourceRelationQueryRequest)
    return nil
}

// ParamYxWarehouseResourceRelationQueryRequest Getter
func (r WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIRequest) GetParamYxWarehouseResourceRelationQueryRequest() *YxWarehouseResourceRelationQueryRequest {
    return r._paramYxWarehouseResourceRelationQueryRequest
}
