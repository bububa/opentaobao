package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订购记录导出 APIResponse
taobao.vas.subsc.search

用于ISV查询自己名下的应用及收费项目的订购记录
*/
type TaobaoVasSubscSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVasSubscSearchResponse `json:"vas_subsc_search_response,omitempty"` 
    TaobaoVasSubscSearchResponse
}

/* model for simplify = false
type TaobaoVasSubscSearchResponse struct {

    // 订购关系对象
    
    ArticleSubs  struct {
        ArticleSub  []ArticleSub `json:"article_sub,omitempty"`
    } `json:"article_subs,omitempty"`
    

    // 总记录数
    
    TotalItem   int64 `json:"total_item,omitempty"`
    

}
*/

type TaobaoVasSubscSearchResponse struct {

    // 订购关系对象
    ArticleSubs   []ArticleSub `json:"article_subs,omitempty"`

    // 总记录数
    TotalItem   int64 `json:"total_item,omitempty"`

}
