package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推荐词-词推词 APIResponse
alibaba.scbp.reckeyword.search

推荐词-词推词
*/
type AlibabaScbpReckeywordSearchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpReckeywordSearchResponse `json:"alibaba_scbp_reckeyword_search_response,omitempty"` 
    AlibabaScbpReckeywordSearchResponse
}

/* model for simplify = false
type AlibabaScbpReckeywordSearchResponse struct {

    // 词推词结果列表
    
    ResultList  struct {
        RecKeywordDto  []RecKeywordDto `json:"rec_keyword_dto,omitempty"`
    } `json:"result_list,omitempty"`
    

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

}
*/

type AlibabaScbpReckeywordSearchResponse struct {

    // 词推词结果列表
    ResultList   []RecKeywordDto `json:"result_list,omitempty"`

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

}
