package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词报告 API返回值 
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告
*/
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel
}

// 单个关键词报告 成功返回结果
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_report_query_single_keyword_effect_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回参数
    Result   *KeywordReportDto `json:"result,omitempty" xml:"result,omitempty"`
}
