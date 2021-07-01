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
type AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest struct {
    model.Params
    // 变更认证回调参数
    _paramChangeAuthCallBackDo   *ChangeAuthCallBackDo
}

// 初始化AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest对象
func NewAlibabaNazcaAuthChangeauthapplyCallbackRequest() *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest{
    return &AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetApiMethodName() string {
    return "alibaba.nazca.auth.changeauthapply.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamChangeAuthCallBackDo Setter
// 变更认证回调参数
func (r *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) SetParamChangeAuthCallBackDo(_paramChangeAuthCallBackDo *ChangeAuthCallBackDo) error {
    r._paramChangeAuthCallBackDo = _paramChangeAuthCallBackDo
    r.Set("param_change_auth_call_back_do", _paramChangeAuthCallBackDo)
    return nil
}

// ParamChangeAuthCallBackDo Getter
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetParamChangeAuthCallBackDo() *ChangeAuthCallBackDo {
    return r._paramChangeAuthCallBackDo
}
