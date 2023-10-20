package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryBidwordAPIResponse 关键词报表查询 API返回值
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
type TaobaoUniversalbpReportQueryBidwordAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryBidwordAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryBidwordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryBidwordAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryBidwordAPIResponseModel is 关键词报表查询 成功返回结果
type TaobaoUniversalbpReportQueryBidwordAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_bidword_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryBidwordTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryBidwordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryBidwordAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryBidwordAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryBidwordAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryBidwordAPIResponse
func GetTaobaoUniversalbpReportQueryBidwordAPIResponse() *TaobaoUniversalbpReportQueryBidwordAPIResponse {
	return poolTaobaoUniversalbpReportQueryBidwordAPIResponse.Get().(*TaobaoUniversalbpReportQueryBidwordAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryBidwordAPIResponse 将 TaobaoUniversalbpReportQueryBidwordAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryBidwordAPIResponse(v *TaobaoUniversalbpReportQueryBidwordAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryBidwordAPIResponse.Put(v)
}
