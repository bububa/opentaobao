package kclub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索(分页) APIResponse
alibaba.kclub.kc.qa.search.page

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchPageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_kclub_kc_qa_search_page_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaKclubKcQaSearchPageResult `json:"result,omitempty" xml:"