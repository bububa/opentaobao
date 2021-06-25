package antifraud

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
人机识别 APIRequest
taobao.collinafacade.nocaptcha.sig.authenticate

人机识别颁发签名串后,本接口负责向ISV提供签名串校验服务
*/
type TaobaoCollinafacadeNocaptchaSigAuthenticateRequest struct {
    model.Params

    // 签名串校验接口入参
    sigAuthenticateContext   *SigAuthenticateContext 

}

func NewTaobaoCollinafacadeNocaptchaSigAuthenticateRequest() *TaobaoCollinafacadeNocaptchaSigAuthenticateRequest{
    return &TaobaoCollinafacadeNocaptchaSigAuthenticateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCollinafacadeNocaptchaSigAuthenticateRequest) GetApiMethodName() string {
    return "taobao.collinafacade.nocaptcha.sig.authenticate"
}

func (r TaobaoCollinafacadeNocaptchaSigAuthenticateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCollinafacadeNocaptchaSigAuthenticateRequest) SetSigAuthenticateContext(sigAuthenticateContext *SigAuthenticateContext) error {
    r.sigAuthenticateContext = sigAuthenticateContext
    r.Set("sig_authenticate_context", sigAuthenticateContext)
    return nil
}

func (r TaobaoCollinafacadeNocaptchaSigAuthenticateRequest) GetSigAuthenticateContext() *SigAuthenticateContext {
    return r.sigAuthenticateContext
}

