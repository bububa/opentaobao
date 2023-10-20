package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportQueryKeywordEffectAPIResponse 关键词报告 API返回值
// alibaba.scbp.ad.report.query.keyword.effect
//
// 关键词报告
type AlibabaScbpAdReportQueryKeywordEffectAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdReportQueryKeywordEffectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportQueryKeywordEffectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdReportQueryKeywordEffectAPIResponseModel).Reset()
}

// AlibabaScbpAdReportQueryKeywordEffectAPIResponseModel is 关键词报告 成功返回结果
type AlibabaScbpAdReportQueryKeywordEffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_query_keyword_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *KeywordReportDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportQueryKeywordEffectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdReportQueryKeywordEffectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdReportQueryKeywordEffectAPIResponse)
	},
}

// GetAlibabaScbpAdReportQueryKeywordEffectAPIResponse 从 sync.Pool 获取 AlibabaScbpAdReportQueryKeywordEffectAPIResponse
func GetAlibabaScbpAdReportQueryKeywordEffectAPIResponse() *AlibabaScbpAdReportQueryKeywordEffectAPIResponse {
	return poolAlibabaScbpAdReportQueryKeywordEffectAPIResponse.Get().(*AlibabaScbpAdReportQueryKeywordEffectAPIResponse)
}

// ReleaseAlibabaScbpAdReportQueryKeywordEffectAPIResponse 将 AlibabaScbpAdReportQueryKeywordEffectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdReportQueryKeywordEffectAPIResponse(v *AlibabaScbpAdReportQueryKeywordEffectAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdReportQueryKeywordEffectAPIResponse.Put(v)
}
