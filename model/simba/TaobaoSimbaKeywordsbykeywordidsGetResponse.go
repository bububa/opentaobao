package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据一个关键词Id列表取得一组关键词 APIResponse
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordsbykeywordidsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordsbykeywordidsGetResponse `json:"simba_keywordsbykeywordids_get_response,omitempty"` 
    TaobaoSimbaKeywordsbykeywordidsGetResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordsbykeywordidsGetResponse struct {

    // 取得的关键词列表
    
    Keywords  struct {
        Keyword  []Keyword `json:"keyword,omitempty"`
    } `json:"keywords,omitempty"`
    

}
*/

type TaobaoSimbaKeywordsbykeywordidsGetResponse struct {

    // 取得的关键词列表
    Keywords   []Keyword `json:"keywords,omitempty"`

}
