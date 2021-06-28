package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划关键词数目 APIResponse
alibaba.scbp.ad.keyword.get.keyword.count.by.query

计划关键词数目
*/
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_get_keyword_count_by_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值
    
    Result   int64 `json:"result,omitempty" xml:"