package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建一批关键词 APIResponse
taobao.simba.keywordsvon.add

创建一批关键词
*/
type TaobaoSimbaKeywordsvonAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordsvonAddResponse `json:"simba_keywordsvon_add_response,omitempty"` 
    TaobaoSimbaKeywordsvonAddResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordsvonAddResponse struct {

    // 关键词列表
    
    Keywords  struct {
        Keyword  []Keyword `json:"keyword,omitempty"`
    } `json:"keywords,omitempty"`
    

}
*/

type TaobaoSimbaKeywordsvonAddResponse struct {

    // 关键词列表
    Keywords   []Keyword `json:"keywords,omitempty"`

}
