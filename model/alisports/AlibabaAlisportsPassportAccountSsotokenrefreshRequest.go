package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sso_token刷新 API请求
alibaba.alisports.passport.account.ssotokenrefresh

sso_token刷新
*/
type AlibabaAlisportsPassportAccountSsotokenrefreshRequest struct {
    model.Params
    // 应用appkey
    alispAppKey   string
    // 当前时间戳[精确到秒,10位]
    alispTime   string
    // 签名
    alispSign   string
    // 登录成功返回的access_token(access_token是TOP保留关键字，只能改名，望谅解)
    secret   string
}

// 初始化AlibabaAlisportsPassportAccountSsotokenrefreshRequest对象
func NewAlibabaAlisportsPassportAccountSsotokenrefreshRequest() *AlibabaAlisportsPassportAccountSsotokenrefreshRequest{
    return &AlibabaAlisportsPassportAccountSsotokenrefreshRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.ssotokenrefresh"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispTime Setter
// 当前时间戳[精确到秒,10位]
func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispSign Setter
// 签名
func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetAlispSign() string {
    return r.alispSign
}
// Secret Setter
// 登录成功返回的access_token(access_token是TOP保留关键字，只能改名，望谅解)
func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetSecret(secret string) error {
    r.secret = secret
    r.Set("secret", secret)
    return nil
}

// Secret Getter
func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetSecret() string {
    return r.secret
}
