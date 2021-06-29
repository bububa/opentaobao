package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
open account token验证 API请求
taobao.open.account.token.validate

open account token验证
*/
type TaobaoOpenAccountTokenValidateRequest struct {
    model.Params
    // token
    paramToken   string
}

// 初始化TaobaoOpenAccountTokenValidateRequest对象
func NewTaobaoOpenAccountTokenValidateRequest() *TaobaoOpenAccountTokenValidateRequest{
    return &TaobaoOpenAccountTokenValidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountTokenValidateRequest) GetApiMethodName() string {
    return "taobao.open.account.token.validate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountTokenValidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamToken Setter
// token
func (r *TaobaoOpenAccountTokenValidateRequest) SetParamToken(paramToken string) error {
    r.paramToken = paramToken
    r.Set("param_token", paramToken)
    return nil
}

// ParamToken Getter
func (r TaobaoOpenAccountTokenValidateRequest) GetParamToken() string {
    return r.paramToken
}
