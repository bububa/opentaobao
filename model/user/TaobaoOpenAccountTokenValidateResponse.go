package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
open account token验证 APIResponse
taobao.open.account.token.validate

open account token验证
*/
type TaobaoOpenAccountTokenValidateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenAccountTokenValidateResponse `json:"open_account_token_validate_response,omitempty"` 
    TaobaoOpenAccountTokenValidateResponse
}

/* model for simplify = false
type TaobaoOpenAccountTokenValidateResponse struct {

    // 验证成功返回token中的信息
    
    Data  *struct {
        OpenAccountTokenValidateResult  *OpenAccountTokenValidateResult `json:"open_account_token_validate_result,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoOpenAccountTokenValidateResponse struct {

    // 验证成功返回token中的信息
    Data   *OpenAccountTokenValidateResult `json:"data,omitempty"`

}
