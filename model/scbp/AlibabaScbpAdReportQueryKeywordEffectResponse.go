package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报告 APIResponse
alibaba.scbp.ad.report.query.keyword.effect

关键词报告
*/
type AlibabaScbpAdReportQueryKeywordEffectAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdReportQueryKeywordEffectResponse
}

type AlibabaScbpAdReportQueryKeywordEffectResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_report_query_keyword_effect_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回数据
    
    Result   *KeywordReportDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
