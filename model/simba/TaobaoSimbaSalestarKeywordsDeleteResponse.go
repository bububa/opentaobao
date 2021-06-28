package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销量明星关键词删除 APIResponse
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口
*/
type TaobaoSimbaSalestarKeywordsDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSalestarKeywordsDeleteResponse `json:"simba_salestar_keywords_delete_response,omitempty"` 
    TaobaoSimbaSalestarKeywordsDeleteResponse
}

/* model for simplify = false
type TaobaoSimbaSalestarKeywordsDeleteResponse struct {

    // 删除成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 成功删除条数
    
    Results   int64 `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaSalestarKeywordsDeleteResponse struct {

    // 删除成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 成功删除条数
    Results   int64 `json:"results,omitempty"`

}
