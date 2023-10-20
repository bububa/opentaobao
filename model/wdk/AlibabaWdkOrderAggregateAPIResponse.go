package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkorderaggregateAPIResponse 淘鲜达订单按门店机台号聚合查询 API返回值
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
type AlibabawdkorderaggregateAPIResponse struct {
	model.CommonResponse
	AlibabawdkorderaggregateAPIResponseModel
}

// AlibabawdkorderaggregateAPIResponseModel is 淘鲜达订单按门店机台号聚合查询 成功返回结果
type AlibabawdkorderaggregateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_order_aggregate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderAggregateQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
