package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse 中心仓查网格仓 API返回值
// wdk.logistic.network.warehouse.resource.relation.query.from
//
// 盒马集市，中心仓查询网格仓
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse struct {
	model.CommonResponse
	WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponseModel
}

// Reset 清空结构体
func (m *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponseModel).Reset()
}

// WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponseModel is 中心仓查网格仓 成功返回结果
type WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_logistic_network_warehouse_resource_relation_query_from_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse)
	},
}

// GetWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse 从 sync.Pool 获取 WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse
func GetWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse() *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse {
	return poolWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse.Get().(*WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse)
}

// ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse 将 WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse 保存到 sync.Pool
func ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse(v *WdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse) {
	v.Reset()
	poolWdkLogisticNetworkWarehouseResourceRelationQueryFromAPIResponse.Put(v)
}
