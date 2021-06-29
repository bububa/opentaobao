package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sso_token刷新 APIRequest
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

func NewAlibabaAlisportsPassportAccountSsotokenrefreshRequest() *AlibabaAlisportsPassportAccountSsotokenrefreshRequest{
    return &AlibabaAlisportsPassportAccountSsotokenrefreshRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.ssotokenrefresh"
}

func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsPassportAccountSsotokenrefreshRequest) SetSecret(secret string) error {
    r.secret = secret
    r.Set("secret", secret)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenrefreshRequest) GetSecret() string {
    return r.secret
}

