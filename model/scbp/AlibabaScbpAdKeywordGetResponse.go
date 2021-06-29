package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
外贸直通车查询关键词 API返回值 
alibaba.scbp.ad.keyword.get

外贸直通车查询关键词
*/
type AlibabaScbpAdKeywordGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdKeywordGetResponse
}

// 外贸直通车查询关键词 成功返回结果
type AlibabaScbpAdKeywordGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_keyword_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 总个数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // 查询关键词列表
    KeywordList   []KeywordResultDTO `json:"keyword_list,omitempty" xml:"keyword_list>keyword_result_dto,omitempty"`
}
