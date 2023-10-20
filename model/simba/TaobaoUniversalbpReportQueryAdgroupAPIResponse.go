package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAdgroupAPIResponse 单元报表查询 API返回值
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
type TaobaoUniversalbpReportQueryAdgroupAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryAdgroupAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryAdgroupAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryAdgroupAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryAdgroupAPIResponseModel is 单元报表查询 成功返回结果
type TaobaoUniversalbpReportQueryAdgroupAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_adgroup_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryAdgroupTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryAdgroupAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryAdgroupAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryAdgroupAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryAdgroupAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryAdgroupAPIResponse
func GetTaobaoUniversalbpReportQueryAdgroupAPIResponse() *TaobaoUniversalbpReportQueryAdgroupAPIResponse {
	return poolTaobaoUniversalbpReportQueryAdgroupAPIResponse.Get().(*TaobaoUniversalbpReportQueryAdgroupAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryAdgroupAPIResponse 将 TaobaoUniversalbpReportQueryAdgroupAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryAdgroupAPIResponse(v *TaobaoUniversalbpReportQueryAdgroupAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryAdgroupAPIResponse.Put(v)
}
