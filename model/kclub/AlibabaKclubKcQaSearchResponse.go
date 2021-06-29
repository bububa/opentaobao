package kclub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-知识检索 APIResponse
alibaba.kclub.kc.qa.search

知识云-知识搜索服务
*/
type AlibabaKclubKcQaSearchAPIResponse struct {
    model.CommonResponse
    AlibabaKclubKcQaSearchResponse
}

type AlibabaKclubKcQaSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_kclub_kc_qa_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 搜索结果
    
    Result   *AlibabaKclubKcQaSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
