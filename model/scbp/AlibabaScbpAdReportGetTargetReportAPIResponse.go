package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetTargetReportAPIResponse 定向报告 API返回值
// alibaba.scbp.ad.report.get.target.report
//
// 定向报告
type AlibabaScbpAdReportGetTargetReportAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdReportGetTargetReportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetTargetReportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdReportGetTargetReportAPIResponseModel).Reset()
}

// AlibabaScbpAdReportGetTargetReportAPIResponseModel is 定向报告 成功返回结果
type AlibabaScbpAdReportGetTargetReportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_get_target_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *TargetReportDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetTargetReportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdReportGetTargetReportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdReportGetTargetReportAPIResponse)
	},
}

// GetAlibabaScbpAdReportGetTargetReportAPIResponse 从 sync.Pool 获取 AlibabaScbpAdReportGetTargetReportAPIResponse
func GetAlibabaScbpAdReportGetTargetReportAPIResponse() *AlibabaScbpAdReportGetTargetReportAPIResponse {
	return poolAlibabaScbpAdReportGetTargetReportAPIResponse.Get().(*AlibabaScbpAdReportGetTargetReportAPIResponse)
}

// ReleaseAlibabaScbpAdReportGetTargetReportAPIResponse 将 AlibabaScbpAdReportGetTargetReportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdReportGetTargetReportAPIResponse(v *AlibabaScbpAdReportGetTargetReportAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdReportGetTargetReportAPIResponse.Put(v)
}
