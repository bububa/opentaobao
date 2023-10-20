package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse 仓站（网格仓自提点）关系查询 API返回值
// wdk.logistic.network.warehouse.delivery.relation.query
//
// 盒马集市，仓站（网格仓自提点）关系查询
type WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse struct {
	model.CommonResponse
	WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponseModel
}

// Reset 清空结构体
func (m *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponseModel).Reset()
}

// WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponseModel is 仓站（网格仓自提点）关系查询 成功返回结果
type WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_logistic_network_warehouse_delivery_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse)
	},
}

// GetWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse 从 sync.Pool 获取 WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse
func GetWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse() *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse {
	return poolWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse.Get().(*WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse)
}

// ReleaseWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse 将 WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse 保存到 sync.Pool
func ReleaseWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse(v *WdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse) {
	v.Reset()
	poolWdkLogisticNetworkWarehouseDeliveryRelationQueryAPIResponse.Put(v)
}
