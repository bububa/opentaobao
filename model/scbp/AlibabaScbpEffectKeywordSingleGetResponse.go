package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
单个关键词效果报表 APIResponse
alibaba.scbp.effect.keyword.single.get

单个关键词效果报表
*/
type AlibabaScbpEffectKeywordSingleGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpEffectKeywordSingleGetResponse `json:"alibaba_scbp_effect_keyword_single_get_response,omitempty"` 
    AlibabaScbpEffectKeywordSingleGetResponse
}

/* model for simplify = false
type AlibabaScbpEffectKeywordSingleGetResponse struct {

    // 单个关键词报表
    
    KeywordEffectList  struct {
        SingleAdKeywordEffectDto  []SingleAdKeywordEffectDto `json:"single_ad_keyword_effect_dto,omitempty"`
    } `json:"keyword_effect_list,omitempty"`
    

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

}
*/

type AlibabaScbpEffectKeywordSingleGetResponse struct {

    // 单个关键词报表
    KeywordEffectList   []SingleAdKeywordEffectDto `json:"keyword_effect_list,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

}
