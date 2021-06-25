package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
计划关键词数目 APIResponse
alibaba.scbp.ad.keyword.get.keyword.count.by.query

计划关键词数目
*/
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdKeywordGetKeywordCountByQueryResponse `json:"alibaba_scbp_ad_keyword_get_keyword_count_by_query_response,omitempty"`
}

type AlibabaScbpAdKeywordGetKeywordCountByQueryResponse struct {

    // 返回值
    Result   int64 `json:"result,omitempty"`

}
