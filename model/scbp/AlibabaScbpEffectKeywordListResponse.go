package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报表 APIResponse
alibaba.scbp.effect.keyword.list

关键词报表
*/
type AlibabaScbpEffectKeywordListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_effect_keyword_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词效果数据列表
    
    KeywordReportList   []AdKeywordEffectDto `json:"keyword_report_list,omitempty" xml:"