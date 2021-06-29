package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓站（网格仓自提点）关系查询 API请求
wdk.logistic.network.warehouse.delivery.relation.query

盒马集市，仓站（网格仓自提点）关系查询
*/
type WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest struct {
    model.Params
    // 参数
    _paramWarehouseDeliveryRelationPageQueryRequest   *WarehouseDeliveryRelationPageQueryRequest
}

// 初始化WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest对象
func NewWdkLogisticNetworkWarehouseDeliveryRelationQueryRequest() *WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest{
    return &WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) GetApiMethodName() string {
    return "wdk.logistic.network.warehouse.delivery.relation.query"
}

// IRequest interface 方法, 获取API参数
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamWarehouseDeliveryRelationPageQueryRequest Setter
// 参数
func (r *WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) SetParamWarehouseDeliveryRelationPageQueryRequest(_paramWarehouseDeliveryRelationPageQueryRequest *WarehouseDeliveryRelationPageQueryRequest) error {
    r._paramWarehouseDeliveryRelationPageQueryRequest = _paramWarehouseDeliveryRelationPageQueryRequest
    r.Set("param_warehouse_delivery_relation_page_query_request", _paramWarehouseDeliveryRelationPageQueryRequest)
    return nil
}

// ParamWarehouseDeliveryRelationPageQueryRequest Getter
func (r WdkLogisticNetworkWarehouseDeliveryRelationQueryRequest) GetParamWarehouseDeliveryRelationPageQueryRequest() *WarehouseDeliveryRelationPageQueryRequest {
    return r._paramWarehouseDeliveryRelationPageQueryRequest
}
