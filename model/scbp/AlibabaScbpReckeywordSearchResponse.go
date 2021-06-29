package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推荐词-词推词 APIResponse
alibaba.scbp.reckeyword.search

推荐词-词推词
*/
type AlibabaScbpReckeywordSearchAPIResponse struct {
    model.CommonResponse
    AlibabaScbpReckeywordSearchResponse
}

type AlibabaScbpReckeywordSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_reckeyword_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 词推词结果列表
    
    ResultList   []RecKeywordDto `json:"result_list,omitempty" xml:"result_list>rec_keyword_dto,omitempty"`
    
    
    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
}
