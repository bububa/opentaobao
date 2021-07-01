package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤匹配 API返回值 
taobao.kfc.keyword.search

对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
*/
type TaobaoKfcKeywordSearchAPIResponse struct {
    model.CommonResponse
    TaobaoKfcKeywordSearchAPIResponseModel
}

// 关键词过滤匹配 成功返回结果
type TaobaoKfcKeywordSearchAPIResponseModel struct {
    XMLName xml.Name `xml:"kfc_keyword_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // KFC 关键词过滤匹配结果
    KfcSearchResult   *KfcSearchResult `json:"kfc_search_result,omitempty" xml:"kfc_search_result,omitempty"`
}
