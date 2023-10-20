package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderUploadReportAPIResponse 服务商上传验货报告同步给闲鱼 API返回值
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
type AlibabaIdleTenderUploadReportAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderUploadReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderUploadReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderUploadReportAPIResponseModel).Reset()
}

// AlibabaIdleTenderUploadReportAPIResponseModel is 服务商上传验货报告同步给闲鱼 成功返回结果
type AlibabaIdleTenderUploadReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_upload_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *IdleCommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderUploadReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleTenderUploadReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderUploadReportAPIResponse)
	},
}

// GetAlibabaIdleTenderUploadReportAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderUploadReportAPIResponse
func GetAlibabaIdleTenderUploadReportAPIResponse() *AlibabaIdleTenderUploadReportAPIResponse {
	return poolAlibabaIdleTenderUploadReportAPIResponse.Get().(*AlibabaIdleTenderUploadReportAPIResponse)
}

// ReleaseAlibabaIdleTenderUploadReportAPIResponse 将 AlibabaIdleTenderUploadReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderUploadReportAPIResponse(v *AlibabaIdleTenderUploadReportAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderUploadReportAPIResponse.Put(v)
}
