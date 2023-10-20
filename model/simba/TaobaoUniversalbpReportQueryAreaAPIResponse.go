package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAreaAPIResponse 地域报表查询 API返回值
// taobao.universalbp.report.query.area
//
// 地域报表查询
type TaobaoUniversalbpReportQueryAreaAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryAreaAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryAreaAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryAreaAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryAreaAPIResponseModel is 地域报表查询 成功返回结果
type TaobaoUniversalbpReportQueryAreaAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_area_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryAreaTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryAreaAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryAreaAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryAreaAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryAreaAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryAreaAPIResponse
func GetTaobaoUniversalbpReportQueryAreaAPIResponse() *TaobaoUniversalbpReportQueryAreaAPIResponse {
	return poolTaobaoUniversalbpReportQueryAreaAPIResponse.Get().(*TaobaoUniversalbpReportQueryAreaAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryAreaAPIResponse 将 TaobaoUniversalbpReportQueryAreaAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryAreaAPIResponse(v *TaobaoUniversalbpReportQueryAreaAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryAreaAPIResponse.Put(v)
}
