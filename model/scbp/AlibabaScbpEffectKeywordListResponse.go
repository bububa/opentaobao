package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词报表 APIResponse
alibaba.scbp.effect.keyword.list

关键词报表
*/
type AlibabaScbpEffectKeywordListAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpEffectKeywordListResponse `json:"alibaba_scbp_effect_keyword_list_response,omitempty"`
}

type AlibabaScbpEffectKeywordListResponse struct {

    // 关键词效果数据列表
    KeywordReportList   []AdKeywordEffectDto `json:"keyword_report_list,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

}
