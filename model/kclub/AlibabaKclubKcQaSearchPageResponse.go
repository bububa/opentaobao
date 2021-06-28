package kclub

import (
    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索(分页) APIResponse
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchPageAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaKclubKcQaSearchPageResponse `json:"alibaba_kclub_kc_qa_search_page_response,omitempty"` 
    AlibabaKclubKcQaSearchPageResponse
}

/* model for simplify = false
type AlibabaKclubKcQaSearchPageResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaKclubKcQaSearchPageResult  *AlibabaKclubKcQaSearchPageResult `json:"alibaba_kclub_kc_qa_search_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaKclubKcQaSearchPageResponse struct {

    // 返回结果
    Result   *AlibabaKclubKcQaSearchPageResult `json:"result,omitempty"`

}
