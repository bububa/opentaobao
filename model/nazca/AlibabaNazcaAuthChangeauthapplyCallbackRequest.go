package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
变更认证回调 APIRequest
alibaba.nazca.auth.changeauthapply.callback

变更认证回调
*/
type AlibabaNazcaAuthChangeauthapplyCallbackRequest struct {
    model.Params

    // 变更认证回调参数
    paramChangeAuthCallBackDo   *ChangeAuthCallBackDo 

}

func NewAlibabaNazcaAuthChangeauthapplyCallbackRequest() *AlibabaNazcaAuthChangeauthapplyCallbackRequest{
    return &AlibabaNazcaAuthChangeauthapplyCallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNazcaAuthChangeauthapplyCallbackRequest) GetApiMethodName() string {
    return "alibaba.nazca.auth.changeauthapply.callback"
}

func (r AlibabaNazcaAuthChangeauthapplyCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNazcaAuthChangeauthapplyCallbackRequest) SetParamChangeAuthCallBackDo(paramChangeAuthCallBackDo *ChangeAuthCallBackDo) error {
    r.paramChangeAuthCallBackDo = paramChangeAuthCallBackDo
    r.Set("param_change_auth_call_back_do", paramChangeAuthCallBackDo)
    return nil
}

func (r AlibabaNazcaAuthChangeauthapplyCallbackRequest) GetParamChangeAuthCallBackDo() *ChangeAuthCallBackDo {
    return r.paramChangeAuthCallBackDo
}

