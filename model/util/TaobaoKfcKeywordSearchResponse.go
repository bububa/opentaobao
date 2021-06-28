package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词过滤匹配 APIResponse
taobao.kfc.keyword.search

对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
*/
type TaobaoKfcKeywordSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoKfcKeywordSearchResponse `json:"kfc_keyword_search_response,omitempty"` 
    TaobaoKfcKeywordSearchResponse
}

/* model for simplify = false
type TaobaoKfcKeywordSearchResponse struct {

    // KFC 关键词过滤匹配结果
    
    KfcSearchResult  *struct {
        KfcSearchResult  *KfcSearchResult `json:"kfc_search_result,omitempty"`
    } `json:"kfc_search_result,omitempty"`
    

}
*/

type TaobaoKfcKeywordSearchResponse struct {

    // KFC 关键词过滤匹配结果
    KfcSearchResult   *KfcSearchResult `json:"kfc_search_result,omitempty"`

}
