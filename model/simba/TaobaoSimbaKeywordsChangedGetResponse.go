package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的关键词ID、宝贝id、修改时间 APIResponse
taobao.simba.keywords.changed.get

分页获取修改过的关键词ID、宝贝id、修改时间
*/
type TaobaoSimbaKeywordsChangedGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordsChangedGetResponse `json:"simba_keywords_changed_get_response,omitempty"` 
    TaobaoSimbaKeywordsChangedGetResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordsChangedGetResponse struct {

    // 关键词分页对象
    
    Keywords  *struct {
        KeywordPage  *KeywordPage `json:"keyword_page,omitempty"`
    } `json:"keywords,omitempty"`
    

}
*/

type TaobaoSimbaKeywordsChangedGetResponse struct {

    // 关键词分页对象
    Keywords   *KeywordPage `json:"keywords,omitempty"`

}
