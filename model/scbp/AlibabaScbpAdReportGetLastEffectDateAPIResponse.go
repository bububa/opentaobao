package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportGetLastEffectDateAPIResponse 获取最近报表生成时间 API返回值
// alibaba.scbp.ad.report.get.last.effect.date
//
// 获取最近报表生成时间
type AlibabaScbpAdReportGetLastEffectDateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdReportGetLastEffectDateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetLastEffectDateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdReportGetLastEffectDateAPIResponseModel).Reset()
}

// AlibabaScbpAdReportGetLastEffectDateAPIResponseModel is 获取最近报表生成时间 成功返回结果
type AlibabaScbpAdReportGetLastEffectDateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_get_last_effect_date_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 报表最近生成时间
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportGetLastEffectDateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaScbpAdReportGetLastEffectDateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdReportGetLastEffectDateAPIResponse)
	},
}

// GetAlibabaScbpAdReportGetLastEffectDateAPIResponse 从 sync.Pool 获取 AlibabaScbpAdReportGetLastEffectDateAPIResponse
func GetAlibabaScbpAdReportGetLastEffectDateAPIResponse() *AlibabaScbpAdReportGetLastEffectDateAPIResponse {
	return poolAlibabaScbpAdReportGetLastEffectDateAPIResponse.Get().(*AlibabaScbpAdReportGetLastEffectDateAPIResponse)
}

// ReleaseAlibabaScbpAdReportGetLastEffectDateAPIResponse 将 AlibabaScbpAdReportGetLastEffectDateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdReportGetLastEffectDateAPIResponse(v *AlibabaScbpAdReportGetLastEffectDateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdReportGetLastEffectDateAPIResponse.Put(v)
}
