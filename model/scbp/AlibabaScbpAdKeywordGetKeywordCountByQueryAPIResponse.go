package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划关键词数目 API返回值 
alibaba.scbp.ad.keyword.get.keyword.count.by.query

计划关键词数目
*/
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel
}

// 计划关键词数目 成功返回结果
type AlibabaScbpAdKeywordGetKeywordCountByQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_get_keyword_count_by_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`
}
