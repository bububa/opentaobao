package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统帐号登录注册token验证接口 APIRequest
alibaba.alisports.passport.account.tokenvalidate

阿里体育会员系统帐号登录注册token验证接口
*/
type AlibabaAlisportsPassportAccountTokenvalidateRequest struct {
    model.Params

    // 业务方appkey
    alispAppKey   string 

    // 签名
    alispSign   string 

    // token
    token   string 

    // 注册用户类型
    userType   int64 

    // 时间戳
    alispTime   string 

    // 一键登录参数
    secret   string 

    // json字符串，传入扩展字段
    extInfo   string 

    // 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
    mtopAppkey   string 

}

func NewAlibabaAlisportsPassportAccountTokenvalidateRequest() *AlibabaAlisportsPassportAccountTokenvalidateRequest{
    return &AlibabaAlisportsPassportAccountTokenvalidateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.tokenvalidate"
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetToken() string {
    return r.token
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetUserType(userType int64) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetUserType() int64 {
    return r.userType
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetSecret(secret string) error {
    r.secret = secret
    r.Set("secret", secret)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetSecret() string {
    return r.secret
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetExtInfo(extInfo string) error {
    r.extInfo = extInfo
    r.Set("ext_info", extInfo)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetExtInfo() string {
    return r.extInfo
}

func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetMtopAppkey(mtopAppkey string) error {
    r.mtopAppkey = mtopAppkey
    r.Set("mtop_appkey", mtopAppkey)
    return nil
}

func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetMtopAppkey() string {
    return r.mtopAppkey
}

