package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取Access Token APIResponse
taobao.top.auth.token.create

用户通过code换获取access_token，https only
*/
type TaobaoTopAuthTokenCreateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTopAuthTokenCreateResponse `json:"top_auth_token_create_response,omitempty"` 
    TaobaoTopAuthTokenCreateResponse
}

/* model for simplify = false
type TaobaoTopAuthTokenCreateResponse struct {

    // 返回的是json信息，和之前调用https://oauth.taobao.com/tac/token https://oauth.alibaba.com/token 换token返回的字段信息一致
    
    TokenResult   string `json:"token_result,omitempty"`
    

}
*/

type TaobaoTopAuthTokenCreateResponse struct {

    // 返回的是json信息，和之前调用https://oauth.taobao.com/tac/token https://oauth.alibaba.com/token 换token返回的字段信息一致
    TokenResult   string `json:"token_result,omitempty"`

}
