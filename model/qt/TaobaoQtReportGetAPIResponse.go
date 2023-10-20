package qt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportGetAPIResponse 查询质检报告 API返回值
// taobao.qt.report.get
//
// 质检报告查询
type TaobaoQtReportGetAPIResponse struct {
	model.CommonResponse
	TaobaoQtReportGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQtReportGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQtReportGetAPIResponseModel).Reset()
}

// TaobaoQtReportGetAPIResponseModel is 查询质检报告 成功返回结果
type TaobaoQtReportGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 质检报告对象
	QtReport *QtReport `json:"qt_report,omitempty" xml:"qt_report,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQtReportGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.QtReport = nil
}

var poolTaobaoQtReportGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQtReportGetAPIResponse)
	},
}

// GetTaobaoQtReportGetAPIResponse 从 sync.Pool 获取 TaobaoQtReportGetAPIResponse
func GetTaobaoQtReportGetAPIResponse() *TaobaoQtReportGetAPIResponse {
	return poolTaobaoQtReportGetAPIResponse.Get().(*TaobaoQtReportGetAPIResponse)
}

// ReleaseTaobaoQtReportGetAPIResponse 将 TaobaoQtReportGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQtReportGetAPIResponse(v *TaobaoQtReportGetAPIResponse) {
	v.Reset()
	poolTaobaoQtReportGetAPIResponse.Put(v)
}
