package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
open account数据搜索 APIResponse
taobao.open.account.search

open account数据搜索
*/
type TaobaoOpenAccountSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenAccountSearchResponse `json:"open_account_search_response,omitempty"` 
    TaobaoOpenAccountSearchResponse
}

/* model for simplify = false
type TaobaoOpenAccountSearchResponse struct {

    // 返回结果
    
    Data  *struct {
        OpenAccountSearchResult  *OpenAccountSearchResult `json:"open_account_search_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoOpenAccountSearchResponse struct {

    // 返回结果
    Data   *OpenAccountSearchResult `json:"data,omitempty"`

}
