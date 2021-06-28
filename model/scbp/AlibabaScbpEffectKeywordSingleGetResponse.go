package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词效果报表 APIResponse
alibaba.scbp.effect.keyword.single.get

单个关键词效果报表
*/
type AlibabaScbpEffectKeywordSingleGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_effect_keyword_single_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 单个关键词报表
    
    KeywordEffectList   []SingleAdKeywordEffectDto `json:"keyword_effect_list,omitempty" xml:"