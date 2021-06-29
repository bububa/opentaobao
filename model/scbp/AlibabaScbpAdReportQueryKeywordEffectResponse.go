package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报告 API返回值 
alibaba.scbp.ad.report.query.keyword.effect

关键词报告
*/
type AlibabaScbpAdReportQueryKeywordEffectAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdReportQueryKeywordEffectResponse
}

// 关键词报告 成功返回结果
type AlibabaScbpAdReportQueryKeywordEffectResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_report_query_keyword_effect_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Result   *KeywordReportDTO `json:"result,omitempty" xml:"result,omitempty"`
}
