package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词报告 APIResponse
alibaba.scbp.ad.report.query.single.keyword.effect

单个关键词报告
*/
type AlibabaScbpAdReportQuerySingleKeywordEffectAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_report_query_single_keyword_effect_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回参数
    
    Result   *KeywordReportDto `json:"result,omitempty" xml:"