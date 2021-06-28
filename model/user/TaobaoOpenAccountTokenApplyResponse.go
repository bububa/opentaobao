package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
申请免登Open Account Token APIResponse
taobao.open.account.token.apply

申请免登Open Account Token
*/
type TaobaoOpenAccountTokenApplyAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenAccountTokenApplyResponse `json:"open_account_token_apply_response,omitempty"` 
    TaobaoOpenAccountTokenApplyResponse
}

/* model for simplify = false
type TaobaoOpenAccountTokenApplyResponse struct {

    // 返回的token结果
    
    Data  *struct {
        OpenAccountTokenApplyResult  *OpenAccountTokenApplyResult `json:"open_account_token_apply_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoOpenAccountTokenApplyResponse struct {

    // 返回的token结果
    Data   *OpenAccountTokenApplyResult `json:"data,omitempty"`

}
