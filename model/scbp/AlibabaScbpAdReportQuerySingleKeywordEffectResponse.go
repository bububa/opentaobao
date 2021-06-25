package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词报告 APIResponse
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告
*/
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdReportQuerySingleKeywordEffectResponse `json:"alibaba_scbp_ad_report_query_single_keyword_effect_response,omitempty"`
}

type AlibabaScbpAdReportQuerySingleKeywordEffectResponse struct {

    // 返回参数
    Result   *KeywordReportDto `json:"result,omitempty"`

}
