package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponse 按网格仓查中心仓（带缓存） API返回值
// wdk.logistic.network.warehouse.resource.relation.query.to.codes
//
// 盒马集市，网格仓查询中心仓
type WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponse struct {
	model.CommonResponse
	WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponseModel
}

// WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponseModel is 按网格仓查中心仓（带缓存） 成功返回结果
type WdklogisticnetworkwarehouseresourcerelationquerytocodesAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_logistic_network_warehouse_resource_relation_query_to_codes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`
}
