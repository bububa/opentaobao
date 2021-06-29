package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询网格仓-区块-自提点关系 APIResponse
wdk.logistic.network.resource.group.query

查询网格仓-区块-自提点关系
*/
type WdkLogisticNetworkResourceGroupQueryAPIResponse struct {
    model.CommonResponse
    WdkLogisticNetworkResourceGroupQueryResponse
}

type WdkLogisticNetworkResourceGroupQueryResponse struct {
    XMLName xml.Name `xml:"wdk_logistic_network_resource_group_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
