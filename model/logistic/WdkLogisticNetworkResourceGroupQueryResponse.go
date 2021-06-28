package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询网格仓-区块-自提点关系 APIResponse
wdk.logistic.network.resource.group.query

查询网格仓-区块-自提点关系
*/
type WdkLogisticNetworkResourceGroupQueryAPIResponse struct {
    model.CommonResponse
    // Response *WdkLogisticNetworkResourceGroupQueryResponse `json:"wdk_logistic_network_resource_group_query_response,omitempty"` 
    WdkLogisticNetworkResourceGroupQueryResponse
}

/* model for simplify = false
type WdkLogisticNetworkResourceGroupQueryResponse struct {

    // 出参
    
    Result  *struct {
        LogisticsResult  *LogisticsResult `json:"logistics_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type WdkLogisticNetworkResourceGroupQueryResponse struct {

    // 出参
    Result   *LogisticsResult `json:"result,omitempty"`

}
