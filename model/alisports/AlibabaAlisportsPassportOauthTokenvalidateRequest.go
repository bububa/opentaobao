package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统--获取登录信息接口 APIRequest
alibaba.alisports.passport.oauth.tokenvalidate

阿里体育会员系统--获取登录信息接口
*/
type AlibabaAlisportsPassportOauthTokenvalidateRequest struct {
    model.Params

    // 登录时返回给前端的token
    token   string 

    // 时间戳
    alispTime   string 

    // 应用的appkey
    alispAppKey   string 

    // 参数加密之后的串
    alispSign   string 

}

func NewAlibabaAlisportsPassportOauthTokenvalidateRequest() *AlibabaAlisportsPassportOauthTokenvalidateRequest{
    return &AlibabaAlisportsPassportOauthTokenvalidateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.oauth.tokenvalidate"
}

func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetToken() string {
    return r.token
}

func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetAlispSign() string {
    return r.alispSign
}

