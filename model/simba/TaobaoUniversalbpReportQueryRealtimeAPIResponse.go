package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryRealtimeAPIResponse 实时报表查询 API返回值
// taobao.universalbp.report.query.realtime
//
// 实时报表查询
type TaobaoUniversalbpReportQueryRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryRealtimeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryRealtimeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryRealtimeAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryRealtimeAPIResponseModel is 实时报表查询 成功返回结果
type TaobaoUniversalbpReportQueryRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryRealtimeTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryRealtimeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryRealtimeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryRealtimeAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryRealtimeAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryRealtimeAPIResponse
func GetTaobaoUniversalbpReportQueryRealtimeAPIResponse() *TaobaoUniversalbpReportQueryRealtimeAPIResponse {
	return poolTaobaoUniversalbpReportQueryRealtimeAPIResponse.Get().(*TaobaoUniversalbpReportQueryRealtimeAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryRealtimeAPIResponse 将 TaobaoUniversalbpReportQueryRealtimeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryRealtimeAPIResponse(v *TaobaoUniversalbpReportQueryRealtimeAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryRealtimeAPIResponse.Put(v)
}
