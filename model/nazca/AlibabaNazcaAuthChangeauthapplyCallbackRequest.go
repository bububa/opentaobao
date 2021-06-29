package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
变更认证回调 API请求
alibaba.nazca.auth.changeauthapply.callback

变更认证回调
*/
type AlibabaNazcaAuthChangeauthapplyCallbackRequest struct {
    model.Params
    // 变更认证回调参数
    paramChangeAuthCallBackDo   *ChangeAuthCallBackDo
}

// 初始化AlibabaNazcaAuthChangeauthapplyCallbackRequest对象
func NewAlibabaNazcaAuthChangeauthapplyCallbackRequest() *AlibabaNazcaAuthChangeauthapplyCallbackRequest{
    return &AlibabaNazcaAuthChangeauthapplyCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaAuthChangeauthapplyCallbackRequest) GetApiMethodName() string {
    return "alibaba.nazca.auth.changeauthapply.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaAuthChangeauthapplyCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamChangeAuthCallBackDo Setter
// 变更认证回调参数
func (r *AlibabaNazcaAuthChangeauthapplyCallbackRequest) SetParamChangeAuthCallBackDo(paramChangeAuthCallBackDo *ChangeAuthCallBackDo) error {
    r.paramChangeAuthCallBackDo = paramChangeAuthCallBackDo
    r.Set("param_change_auth_call_back_do", paramChangeAuthCallBackDo)
    return nil
}

// ParamChangeAuthCallBackDo Getter
func (r AlibabaNazcaAuthChangeauthapplyCallbackRequest) GetParamChangeAuthCallBackDo() *ChangeAuthCallBackDo {
    return r.paramChangeAuthCallBackDo
}
