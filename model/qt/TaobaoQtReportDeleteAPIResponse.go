package qt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportDeleteAPIResponse 质检报告删除接口 API返回值
// taobao.qt.report.delete
//
// 删除质检报告
type TaobaoQtReportDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoQtReportDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQtReportDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQtReportDeleteAPIResponseModel).Reset()
}

// TaobaoQtReportDeleteAPIResponseModel is 质检报告删除接口 成功返回结果
type TaobaoQtReportDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQtReportDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoQtReportDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQtReportDeleteAPIResponse)
	},
}

// GetTaobaoQtReportDeleteAPIResponse 从 sync.Pool 获取 TaobaoQtReportDeleteAPIResponse
func GetTaobaoQtReportDeleteAPIResponse() *TaobaoQtReportDeleteAPIResponse {
	return poolTaobaoQtReportDeleteAPIResponse.Get().(*TaobaoQtReportDeleteAPIResponse)
}

// ReleaseTaobaoQtReportDeleteAPIResponse 将 TaobaoQtReportDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQtReportDeleteAPIResponse(v *TaobaoQtReportDeleteAPIResponse) {
	v.Reset()
	poolTaobaoQtReportDeleteAPIResponse.Put(v)
}
