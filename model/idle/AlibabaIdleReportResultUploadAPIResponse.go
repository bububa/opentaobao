package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleReportResultUploadAPIResponse 服务商上传验货报告 API返回值
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
type AlibabaIdleReportResultUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleReportResultUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleReportResultUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleReportResultUploadAPIResponseModel).Reset()
}

// AlibabaIdleReportResultUploadAPIResponseModel is 服务商上传验货报告 成功返回结果
type AlibabaIdleReportResultUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_report_result_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleReportResultUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleReportResultUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleReportResultUploadAPIResponse)
	},
}

// GetAlibabaIdleReportResultUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleReportResultUploadAPIResponse
func GetAlibabaIdleReportResultUploadAPIResponse() *AlibabaIdleReportResultUploadAPIResponse {
	return poolAlibabaIdleReportResultUploadAPIResponse.Get().(*AlibabaIdleReportResultUploadAPIResponse)
}

// ReleaseAlibabaIdleReportResultUploadAPIResponse 将 AlibabaIdleReportResultUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleReportResultUploadAPIResponse(v *AlibabaIdleReportResultUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleReportResultUploadAPIResponse.Put(v)
}
