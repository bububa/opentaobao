package antifraud

import (
    "github.com/bububa/opentaobao/model"
)

/* 
人机识别 APIResponse
taobao.collinafacade.nocaptcha.sig.authenticate

人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务
*/
type TaobaoCollinafacadeNocaptchaSigAuthenticateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCollinafacadeNocaptchaSigAuthenticateResponse `json:"taobao_collinafacade_nocaptcha_sig_authenticate_response,omitempty"`
}

type TaobaoCollinafacadeNocaptchaSigAuthenticateResponse struct {

    // 服务出参
    Ret   int64 `json:"ret,omitempty"`

    // 返回authenticateResult
    RetDetail   *SigAuthenticateResult `json:"ret_detail,omitempty"`

}
