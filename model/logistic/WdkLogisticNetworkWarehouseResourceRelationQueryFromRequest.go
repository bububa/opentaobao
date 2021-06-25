package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中心仓查网格仓 APIRequest
wdk.logistic.network.warehouse.resource.relation.query.from

盒马集市，中心仓查询网格仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest struct {
    model.Params

    // 查询参数
    paramPageQueryWarehouseResourceRelationByFromRequest   *PageQueryWarehouseResourceRelationByFromRequest 

}

func NewWdkLogisticNetworkWarehouseResourceRelationQueryFromRequest() *WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest{
    return &WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest{
        Params: model.NewParams(),
    }
}

func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest) GetApiMethodName() string {
    return "wdk.logistic.network.warehouse.resource.relation.query.from"
}

func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest) SetParamPageQueryWarehouseResourceRelationByFromRequest(paramPageQueryWarehouseResourceRelationByFromRequest *PageQueryWarehouseResourceRelationByFromRequest) error {
    r.paramPageQueryWarehouseResourceRelationByFromRequest = paramPageQueryWarehouseResourceRelationByFromRequest
    r.Set("param_page_query_warehouse_resource_relation_by_from_request", paramPageQueryWarehouseResourceRelationByFromRequest)
    return nil
}

func (r WdkLogisticNetworkWarehouseResourceRelationQueryFromRequest) GetParamPageQueryWarehouseResourceRelationByFromRequest() *PageQueryWarehouseResourceRelationByFromRequest {
    return r.paramPageQueryWarehouseResourceRelationByFromRequest
}

