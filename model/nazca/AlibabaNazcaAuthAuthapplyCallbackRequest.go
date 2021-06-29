package nazca

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
认证的统一回调接口 API请求
alibaba.nazca.auth.authapply.callback

认证的统一回调接口
*/
type AlibabaNazcaAuthAuthapplyCallbackRequest struct {
    model.Params
    // 认证回调参数
    authApplyDoneCallbackDo   *AuthApplyDoneCallBackDo
}

// 初始化AlibabaNazcaAuthAuthapplyCallbackRequest对象
func NewAlibabaNazcaAuthAuthapplyCallbackRequest() *AlibabaNazcaAuthAuthapplyCallbackRequest{
    return &AlibabaNazcaAuthAuthapplyCallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNazcaAuthAuthapplyCallbackRequest) GetApiMethodName() string {
    return "alibaba.nazca.auth.authapply.callback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNazcaAuthAuthapplyCallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuthApplyDoneCallbackDo Setter
// 认证回调参数
func (r *AlibabaNazcaAuthAuthapplyCallbackRequest) SetAuthApplyDoneCallbackDo(authApplyDoneCallbackDo *AuthApplyDoneCallBackDo) error {
    r.authApplyDoneCallbackDo = authApplyDoneCallbackDo
    r.Set("auth_apply_done_callback_do", authApplyDoneCallbackDo)
    return nil
}

// AuthApplyDoneCallbackDo Getter
func (r AlibabaNazcaAuthAuthapplyCallbackRequest) GetAuthApplyDoneCallbackDo() *AuthApplyDoneCallBackDo {
    return r.authApplyDoneCallbackDo
}
