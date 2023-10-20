package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdklogisticnetworkresourcegroupqueryAPIResponse 查询网格仓-区块-自提点关系 API返回值
// wdk.logistic.network.resource.group.query
//
// 查询网格仓-区块-自提点关系
type WdklogisticnetworkresourcegroupqueryAPIResponse struct {
	model.CommonResponse
	WdklogisticnetworkresourcegroupqueryAPIResponseModel
}

// WdklogisticnetworkresourcegroupqueryAPIResponseModel is 查询网格仓-区块-自提点关系 成功返回结果
type WdklogisticnetworkresourcegroupqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_logistic_network_resource_group_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}
