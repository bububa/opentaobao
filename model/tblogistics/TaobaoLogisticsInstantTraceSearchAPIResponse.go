package tblogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsInstantTraceSearchAPIResponse 物流详情查询 API返回值
// taobao.logistics.instant.trace.search
//
// 物流详情查询
type TaobaoLogisticsInstantTraceSearchAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsInstantTraceSearchAPIResponseModel
}

// TaobaoLogisticsInstantTraceSearchAPIResponseModel is 物流详情查询 成功返回结果
type TaobaoLogisticsInstantTraceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_instant_trace_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoLogisticsInstantTraceSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
