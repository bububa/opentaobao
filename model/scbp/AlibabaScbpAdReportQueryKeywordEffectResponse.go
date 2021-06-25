package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词报告 APIResponse
alibaba.scbp.ad.report.query.keyword.effect

关键词报告
*/
type AlibabaScbpAdReportQueryKeywordEffectAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdReportQueryKeywordEffectResponse `json:"alibaba_scbp_ad_report_query_keyword_effect_response,omitempty"`
}

type AlibabaScbpAdReportQueryKeywordEffectResponse struct {

    // 返回数据
    Result   *KeywordReportDto `json:"result,omitempty"`

}
