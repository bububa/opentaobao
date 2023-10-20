package qt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportAddAPIResponse 上传质检报告 API返回值
// taobao.qt.report.add
//
// 上传质检报告
type TaobaoQtReportAddAPIResponse struct {
	model.CommonResponse
	TaobaoQtReportAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQtReportAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQtReportAddAPIResponseModel).Reset()
}

// TaobaoQtReportAddAPIResponseModel is 上传质检报告 成功返回结果
type TaobaoQtReportAddAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQtReportAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoQtReportAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQtReportAddAPIResponse)
	},
}

// GetTaobaoQtReportAddAPIResponse 从 sync.Pool 获取 TaobaoQtReportAddAPIResponse
func GetTaobaoQtReportAddAPIResponse() *TaobaoQtReportAddAPIResponse {
	return poolTaobaoQtReportAddAPIResponse.Get().(*TaobaoQtReportAddAPIResponse)
}

// ReleaseTaobaoQtReportAddAPIResponse 将 TaobaoQtReportAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQtReportAddAPIResponse(v *TaobaoQtReportAddAPIResponse) {
	v.Reset()
	poolTaobaoQtReportAddAPIResponse.Put(v)
}
