package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCrowdAPIResponse 人群报表查询 API返回值
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
type TaobaoUniversalbpReportQueryCrowdAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryCrowdAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryCrowdAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryCrowdAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryCrowdAPIResponseModel is 人群报表查询 成功返回结果
type TaobaoUniversalbpReportQueryCrowdAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_crowd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryCrowdTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryCrowdAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryCrowdAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryCrowdAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryCrowdAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryCrowdAPIResponse
func GetTaobaoUniversalbpReportQueryCrowdAPIResponse() *TaobaoUniversalbpReportQueryCrowdAPIResponse {
	return poolTaobaoUniversalbpReportQueryCrowdAPIResponse.Get().(*TaobaoUniversalbpReportQueryCrowdAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryCrowdAPIResponse 将 TaobaoUniversalbpReportQueryCrowdAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryCrowdAPIResponse(v *TaobaoUniversalbpReportQueryCrowdAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryCrowdAPIResponse.Put(v)
}
