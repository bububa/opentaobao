package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓站（网格仓自提点）关系查询 APIResponse
wdk.logistic.network.warehouse.delivery.relation.query

盒马集市，仓站（网格仓自提点）关系查询
*/
type WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse struct {
    model.CommonResponse
    // Response *WdkLogisticNetworkWarehouseDeliveryRelationQueryResponse `json:"wdk_logistic_network_warehouse_delivery_relation_query_response,omitempty"` 
    WdkLogisticNetworkWarehouseDeliveryRelationQueryResponse
}

/* model for simplify = false
type WdkLogisticNetworkWarehouseDeliveryRelationQueryResponse struct {

    // 参数
    
    Result  *struct {
        LogisticsResult  *LogisticsResult `json:"logistics_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type WdkLogisticNetworkWarehouseDeliveryRelationQueryResponse struct {

    // 参数
    Result   *LogisticsResult `json:"result,omitempty"`

}
