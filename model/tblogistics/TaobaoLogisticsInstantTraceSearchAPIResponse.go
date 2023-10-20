package tblogistics

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoLogisticsInstantTraceSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsInstantTraceSearchAPIResponseModel).Reset()
}

// TaobaoLogisticsInstantTraceSearchAPIResponseModel is 物流详情查询 成功返回结果
type TaobaoLogisticsInstantTraceSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_instant_trace_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoLogisticsInstantTraceSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsInstantTraceSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoLogisticsInstantTraceSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsInstantTraceSearchAPIResponse)
	},
}

// GetTaobaoLogisticsInstantTraceSearchAPIResponse 从 sync.Pool 获取 TaobaoLogisticsInstantTraceSearchAPIResponse
func GetTaobaoLogisticsInstantTraceSearchAPIResponse() *TaobaoLogisticsInstantTraceSearchAPIResponse {
	return poolTaobaoLogisticsInstantTraceSearchAPIResponse.Get().(*TaobaoLogisticsInstantTraceSearchAPIResponse)
}

// ReleaseTaobaoLogisticsInstantTraceSearchAPIResponse 将 TaobaoLogisticsInstantTraceSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsInstantTraceSearchAPIResponse(v *TaobaoLogisticsInstantTraceSearchAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsInstantTraceSearchAPIResponse.Put(v)
}
