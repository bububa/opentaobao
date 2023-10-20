package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleReportMediaUploadAPIResponse 验货报告上传文件 API返回值
// alibaba.idle.report.media.upload
//
// 服务商上传文件、图片
type AlibabaIdleReportMediaUploadAPIResponse struct {
	model.CommonResponse
	AlibabaIdleReportMediaUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleReportMediaUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleReportMediaUploadAPIResponseModel).Reset()
}

// AlibabaIdleReportMediaUploadAPIResponseModel is 验货报告上传文件 成功返回结果
type AlibabaIdleReportMediaUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_report_media_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleReportMediaUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleReportMediaUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleReportMediaUploadAPIResponse)
	},
}

// GetAlibabaIdleReportMediaUploadAPIResponse 从 sync.Pool 获取 AlibabaIdleReportMediaUploadAPIResponse
func GetAlibabaIdleReportMediaUploadAPIResponse() *AlibabaIdleReportMediaUploadAPIResponse {
	return poolAlibabaIdleReportMediaUploadAPIResponse.Get().(*AlibabaIdleReportMediaUploadAPIResponse)
}

// ReleaseAlibabaIdleReportMediaUploadAPIResponse 将 AlibabaIdleReportMediaUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleReportMediaUploadAPIResponse(v *AlibabaIdleReportMediaUploadAPIResponse) {
	v.Reset()
	poolAlibabaIdleReportMediaUploadAPIResponse.Put(v)
}
