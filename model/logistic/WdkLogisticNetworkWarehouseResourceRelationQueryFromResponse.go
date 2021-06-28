package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
中心仓查网格仓 APIResponse
wdk.logistic.network.warehouse.resource.relation.query.from

盒马集市，中心仓查询网格仓
*/
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse struct {
    model.CommonResponse
    WdkLogisticNetworkWarehouseResourceRelationQueryFromResponse
}

type WdkLogisticNetworkWarehouseResourceRelationQueryFromResponse struct {
    XMLName xml.Name `xml:"wdk_logistic_network_warehouse_resource_relation_query_from_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
