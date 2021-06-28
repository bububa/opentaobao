package kclub

import (
    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索 APIResponse
alibaba.kclub.kc.qa.search

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaKclubKcQaSearchResponse `json:"alibaba_kclub_kc_qa_search_response,omitempty"` 
    AlibabaKclubKcQaSearchResponse
}

/* model for simplify = false
type AlibabaKclubKcQaSearchResponse struct {

    // 搜索结果
    
    Result  *struct {
        AlibabaKclubKcQaSearchResult  *AlibabaKclubKcQaSearchResult `json:"alibaba_kclub_kc_qa_search_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaKclubKcQaSearchResponse struct {

    // 搜索结果
    Result   *AlibabaKclubKcQaSearchResult `json:"result,omitempty"`

}
