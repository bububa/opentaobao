package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sso_token验证 API请求
alibaba.alisports.passport.account.ssotokenvalidate

sso_token验证
*/
type AlibabaAlisportsPassportAccountSsotokenvalidateRequest struct {
    model.Params
    // sso_token
    _ssoToken   string
    // 应用APPKEY
    _alispAppKey   string
    // 当前时间戳[精确到秒，10位]
    _alispTime   string
    // 签名
    _alispSign   string
}

// 初始化AlibabaAlisportsPassportAccountSsotokenvalidateRequest对象
func NewAlibabaAlisportsPassportAccountSsotokenvalidateRequest() *AlibabaAlisportsPassportAccountSsotokenvalidateRequest{
    return &AlibabaAlisportsPassportAccountSsotokenvalidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.ssotokenvalidate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SsoToken Setter
// sso_token
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetSsoToken(_ssoToken string) error {
    r._ssoToken = _ssoToken
    r.Set("sso_token", _ssoToken)
    return nil
}

// SsoToken Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetSsoToken() string {
    return r._ssoToken
}
// AlispAppKey Setter
// 应用APPKEY
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispTime Setter
// 当前时间戳[精确到秒，10位]
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispSign Setter
// 签名
func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetAlispSign() string {
    return r._alispSign
}
