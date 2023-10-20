package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleInspectionReportAPIResponse 鉴定报告 API返回值
// alibaba.idle.recycle.inspection.report
//
// 回收商鉴定报告
type AlibabaIdleRecycleInspectionReportAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleInspectionReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleInspectionReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRecycleInspectionReportAPIResponseModel).Reset()
}

// AlibabaIdleRecycleInspectionReportAPIResponseModel is 鉴定报告 成功返回结果
type AlibabaIdleRecycleInspectionReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_inspection_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RecycleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRecycleInspectionReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRecycleInspectionReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRecycleInspectionReportAPIResponse)
	},
}

// GetAlibabaIdleRecycleInspectionReportAPIResponse 从 sync.Pool 获取 AlibabaIdleRecycleInspectionReportAPIResponse
func GetAlibabaIdleRecycleInspectionReportAPIResponse() *AlibabaIdleRecycleInspectionReportAPIResponse {
	return poolAlibabaIdleRecycleInspectionReportAPIResponse.Get().(*AlibabaIdleRecycleInspectionReportAPIResponse)
}

// ReleaseAlibabaIdleRecycleInspectionReportAPIResponse 将 AlibabaIdleRecycleInspectionReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRecycleInspectionReportAPIResponse(v *AlibabaIdleRecycleInspectionReportAPIResponse) {
	v.Reset()
	poolAlibabaIdleRecycleInspectionReportAPIResponse.Put(v)
}
