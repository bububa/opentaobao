package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkLogisticNetworkResourceGroupQueryAPIResponse 查询网格仓-区块-自提点关系 API返回值
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
type WdkLogisticNetworkResourceGroupQueryAPIResponse struct {
	model.CommonResponse
	WdkLogisticNetworkResourceGroupQueryAPIResponseModel
}

// Reset 清空结构体
func (m *WdkLogisticNetworkResourceGroupQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkLogisticNetworkResourceGroupQueryAPIResponseModel).Reset()
}

// WdkLogisticNetworkResourceGroupQueryAPIResponseModel is 查询网格仓-区块-自提点关系 成功返回结果
type WdkLogisticNetworkResourceGroupQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_logistic_network_resource_group_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkLogisticNetworkResourceGroupQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkLogisticNetworkResourceGroupQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkLogisticNetworkResourceGroupQueryAPIResponse)
	},
}

// GetWdkLogisticNetworkResourceGroupQueryAPIResponse 从 sync.Pool 获取 WdkLogisticNetworkResourceGroupQueryAPIResponse
func GetWdkLogisticNetworkResourceGroupQueryAPIResponse() *WdkLogisticNetworkResourceGroupQueryAPIResponse {
	return poolWdkLogisticNetworkResourceGroupQueryAPIResponse.Get().(*WdkLogisticNetworkResourceGroupQueryAPIResponse)
}

// ReleaseWdkLogisticNetworkResourceGroupQueryAPIResponse 将 WdkLogisticNetworkResourceGroupQueryAPIResponse 保存到 sync.Pool
func ReleaseWdkLogisticNetworkResourceGroupQueryAPIResponse(v *WdkLogisticNetworkResourceGroupQueryAPIResponse) {
	v.Reset()
	poolWdkLogisticNetworkResourceGroupQueryAPIResponse.Put(v)
}
