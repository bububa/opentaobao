package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车查询关键词 APIResponse
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词
*/
type AlibabaScbpAdKeywordGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordGetResponse `json:"alibaba_scbp_ad_keyword_get_response,omitempty"`
}

type AlibabaScbpAdKeywordGetResponse struct {

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 查询关键词列表
    KeywordList   []KeywordResultDto `json:"keyword_list,omitempty"`

}
