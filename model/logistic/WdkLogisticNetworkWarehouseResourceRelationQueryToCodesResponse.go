package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
按网格仓查中心仓（带缓存） APIResponse
wdk.logistic.network.warehouse.resource.relation.query.to.codes

盒马集市，网格仓查询中心仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse struct {
    model.CommonResponse
    // Response *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesResponse `json:"wdk_logistic_network_warehouse_resource_relation_query_to_codes_response,omitempty"` 
    WdkLogisticNetworkWarehouseResourceRelationQueryToCodesResponse
}

/* model for simplify = false
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesResponse struct {

    // 返回值
    
    Result  *struct {
        LogisticsResult  *LogisticsResult `json:"logistics_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesResponse struct {

    // 返回值
    Result   *LogisticsResult `json:"result,omitempty"`

}
