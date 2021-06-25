package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓站（网格仓自提点）关系查询 APIRequest
wdk.logistic.network.warehouse.delivery.relation.query

盒马集市，仓站（网格仓自提点）关系查询
*/
type WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest struct {
    model.Params

    // 参数
    paramWarehouseDeliveryRelationPageQueryRequest   *WarehouseDeliveryRelationPageQueryRequest 

}

func NewWdkLogisticNetworkWarehouseDeliveryRelationQueryRequest() *WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest{
    return &WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest{
        Params: model.NewParams(),
    }
}

func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) GetApiMethodName() string {
    return "wdk.logistic.network.warehouse.delivery.relation.query"
}

func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) SetParamWarehouseDeliveryRelationPageQueryRequest(paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest) error {
    r.paramWarehouseDeliveryRelationPageQueryRequest = paramWarehouseDeliveryRelationPageQueryRequest
    r.Set("param_warehouse_delivery_relation_page_query_request", paramWarehouseDeliveryRelationPageQueryRequest)
    return nil
}

func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) GetParamWarehouseDeliveryRelationPageQueryRequest() *WarehouseDeliveryRelationPageQueryRequest {
    return r.paramWarehouseDeliveryRelationPageQueryRequest
}

