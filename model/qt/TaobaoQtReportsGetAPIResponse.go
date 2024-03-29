package qt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportsGetAPIResponse 批量查询质检报告 API返回值
// taobao.qt.reports.get
//
// 批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
type TaobaoQtReportsGetAPIResponse struct {
	model.CommonResponse
	TaobaoQtReportsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQtReportsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQtReportsGetAPIResponseModel).Reset()
}

// TaobaoQtReportsGetAPIResponseModel is 批量查询质检报告 成功返回结果
type TaobaoQtReportsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_reports_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 质检报告列表
	Reports []QtReport `json:"reports,omitempty" xml:"reports>qt_report,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQtReportsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Reports = m.Reports[:0]
}

var poolTaobaoQtReportsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQtReportsGetAPIResponse)
	},
}

// GetTaobaoQtReportsGetAPIResponse 从 sync.Pool 获取 TaobaoQtReportsGetAPIResponse
func GetTaobaoQtReportsGetAPIResponse() *TaobaoQtReportsGetAPIResponse {
	return poolTaobaoQtReportsGetAPIResponse.Get().(*TaobaoQtReportsGetAPIResponse)
}

// ReleaseTaobaoQtReportsGetAPIResponse 将 TaobaoQtReportsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQtReportsGetAPIResponse(v *TaobaoQtReportsGetAPIResponse) {
	v.Reset()
	poolTaobaoQtReportsGetAPIResponse.Put(v)
}
