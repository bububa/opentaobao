package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
sso_token验证 APIRequest
alibaba.alisports.passport.account.ssotokenvalidate

sso_token验证
*/
type AlibabaAlisportsPassportAccountSsotokenvalidateRequest struct {
    model.Params

    // sso_token
    ssoToken   string 

    // 应用APPKEY
    alispAppKey   string 

    // 当前时间戳[精确到秒，10位]
    alispTime   string 

    // 签名
    alispSign   string 

}

func NewAlibabaAlisportsPassportAccountSsotokenvalidateRequest() *AlibabaAlisportsPassportAccountSsotokenvalidateRequest{
    return &AlibabaAlisportsPassportAccountSsotokenvalidateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.ssotokenvalidate"
}

func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetSsoToken(ssoToken string) error {
    r.ssoToken = ssoToken
    r.Set("sso_token", ssoToken)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetSsoToken() string {
    return r.ssoToken
}

func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportAccountSsotokenvalidateRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportAccountSsotokenvalidateRequest) GetAlispSign() string {
    return r.alispSign
}

