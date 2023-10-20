package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse 单个关键词报告 API返回值
// alibaba.scbp.ad.report.query.single.keyword.effect
//
// 单个关键词报告
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel).Reset()
}

// AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel is 单个关键词报告 成功返回结果
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_report_query_single_keyword_effect_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *KeywordReportDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse)
	},
}

// GetAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse 从 sync.Pool 获取 AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse
func GetAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse() *AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse {
	return poolAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse.Get().(*AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse)
}

// ReleaseAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse 将 AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse(v *AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse.Put(v)
}
