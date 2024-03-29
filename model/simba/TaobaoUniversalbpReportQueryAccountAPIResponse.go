package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAccountAPIResponse 账户报表查询 API返回值
// taobao.universalbp.report.query.account
//
// 账户报表查询
type TaobaoUniversalbpReportQueryAccountAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryAccountAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryAccountAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryAccountAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryAccountAPIResponseModel is 账户报表查询 成功返回结果
type TaobaoUniversalbpReportQueryAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryAccountTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryAccountAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryAccountAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryAccountAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryAccountAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryAccountAPIResponse
func GetTaobaoUniversalbpReportQueryAccountAPIResponse() *TaobaoUniversalbpReportQueryAccountAPIResponse {
	return poolTaobaoUniversalbpReportQueryAccountAPIResponse.Get().(*TaobaoUniversalbpReportQueryAccountAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryAccountAPIResponse 将 TaobaoUniversalbpReportQueryAccountAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryAccountAPIResponse(v *TaobaoUniversalbpReportQueryAccountAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryAccountAPIResponse.Put(v)
}
