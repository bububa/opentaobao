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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_reckeyword_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 词推词结果列表
    
    ResultList   []RecKeywordDto `json:"result_list,omitempty" xml:"