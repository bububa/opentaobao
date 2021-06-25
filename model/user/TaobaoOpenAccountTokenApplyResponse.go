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
    Response *TaobaoOpenAccountTokenApplyResponse `json:"taobao_open_account_token_apply_response,omitempty"`
}

type TaobaoOpenAccountTokenApplyResponse struct {

    // 返回的token结果
    Data   *OpenAccountTokenApplyResult `json:"data,omitempty"`

}
