package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
open account token验证 APIRequest
taobao.open.account.token.validate

open account token验证
*/
type TaobaoOpenAccountTokenValidateRequest struct {
    model.Params

    // token
    paramToken   string 

}

func NewTaobaoOpenAccountTokenValidateRequest() *TaobaoOpenAccountTokenValidateRequest{
    return &TaobaoOpenAccountTokenValidateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountTokenValidateRequest) GetApiMethodName() string {
    return "taobao.open.account.token.validate"
}

func (r TaobaoOpenAccountTokenValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountTokenValidateRequest) SetParamToken(paramToken string) error {
    r.paramToken = paramToken
    r.Set("param_token", paramToken)
    return nil
}

func (r TaobaoOpenAccountTokenValidateRequest) GetParamToken() string {
    return r.paramToken
}

