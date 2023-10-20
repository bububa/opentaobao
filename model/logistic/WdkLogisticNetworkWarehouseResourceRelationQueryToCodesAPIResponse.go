package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse 按网格仓查中心仓（带缓存） API返回值
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse struct {
	model.CommonResponse
	WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponseModel
}

// Reset 清空结构体
func (m *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponseModel).Reset()
}

// WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponseModel is 按网格仓查中心仓（带缓存） 成功返回结果
type WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_logistic_network_warehouse_resource_relation_query_to_codes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse)
	},
}

// GetWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse 从 sync.Pool 获取 WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse
func GetWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse() *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse {
	return poolWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse.Get().(*WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse)
}

// ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse 将 WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse 保存到 sync.Pool
func ReleaseWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse(v *WdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse) {
	v.Reset()
	poolWdkLogisticNetworkWarehouseResourceRelationQueryToCodesAPIResponse.Put(v)
}
