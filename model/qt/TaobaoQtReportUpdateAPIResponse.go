package qt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportUpdateAPIResponse 更新质检报告 API返回值
// taobao.qt.report.update
//
// 更新质检报告
type TaobaoQtReportUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoQtReportUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQtReportUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQtReportUpdateAPIResponseModel).Reset()
}

// TaobaoQtReportUpdateAPIResponseModel is 更新质检报告 成功返回结果
type TaobaoQtReportUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQtReportUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoQtReportUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQtReportUpdateAPIResponse)
	},
}

// GetTaobaoQtReportUpdateAPIResponse 从 sync.Pool 获取 TaobaoQtReportUpdateAPIResponse
func GetTaobaoQtReportUpdateAPIResponse() *TaobaoQtReportUpdateAPIResponse {
	return poolTaobaoQtReportUpdateAPIResponse.Get().(*TaobaoQtReportUpdateAPIResponse)
}

// ReleaseTaobaoQtReportUpdateAPIResponse 将 TaobaoQtReportUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQtReportUpdateAPIResponse(v *TaobaoQtReportUpdateAPIResponse) {
	v.Reset()
	poolTaobaoQtReportUpdateAPIResponse.Put(v)
}
