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
    // Response *AlibabaScbpAdReportQueryKeywordEffectResponse `json:"alibaba_scbp_ad_report_query_keyword_effect_response,omitempty"` 
    AlibabaScbpAdReportQueryKeywordEffectResponse
}

/* model for simplify = false
type AlibabaScbpAdReportQueryKeywordEffectResponse struct {

    // 返回数据
    
    Result  *struct {
        KeywordReportDto  *KeywordReportDto `json:"keyword_report_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaScbpAdReportQueryKeywordEffectResponse struct {

    // 返回数据
    Result   *KeywordReportDto `json:"result,omitempty"`

}
