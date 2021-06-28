package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按网格仓查中心仓（带缓存） APIResponse
wdk.logistic.network.warehouse.resource.relation.query.to.codes

盒马集市，网格仓查询中心仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_logistic_network_warehouse_resource_relation_query_to_codes_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *LogisticsResult `json:"result,omitempty" xml:"