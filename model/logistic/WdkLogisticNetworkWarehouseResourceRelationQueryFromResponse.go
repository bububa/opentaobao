package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
中心仓查网格仓 APIResponse
wdk.logistic.network.warehouse.resource.relation.query.from

盒马集市，中心仓查询网格仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse struct {
    model.CommonResponse
    Response *WdkLogisticNetworkWarehouseResourceRelationQueryFromResponse `json:"wdk_logistic_network_warehouse_resource_relation_query_from_response,omitempty"`
}

type WdkLogisticNetworkWarehouseResourceRelationQueryFromResponse struct {

    // 返回值
    Result   *LogisticsResult `json:"result,omitempty"`

}
