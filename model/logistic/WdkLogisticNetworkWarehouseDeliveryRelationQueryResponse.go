package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓站（网格仓自提点）关系查询 API返回值 
wdk.logistic.network.warehouse.delivery.relation.query

盒马集市，仓站（网格仓自提点）关系查询
*/
type WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse struct {
    model.CommonResponse
    WdkLogisticNetworkWarehouseDeliveryRelationQueryResponse
}

// 仓站（网格仓自提点）关系查询 成功返回结果
type WdkLogisticNetworkWarehouseDeliveryRelationQueryResponse struct {
    XMLName xml.Name `xml:"wdk_logistic_network_warehouse_delivery_relation_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 参数
    Result   *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}
