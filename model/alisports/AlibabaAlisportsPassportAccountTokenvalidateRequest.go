package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统帐号登录注册token验证接口 API请求
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

// 初始化AlibabaAlisportsPassportAccountTokenvalidateRequest对象
func NewAlibabaAlisportsPassportAccountTokenvalidateRequest() *AlibabaAlisportsPassportAccountTokenvalidateRequest{
    return &AlibabaAlisportsPassportAccountTokenvalidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.tokenvalidate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 业务方appkey
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispSign Setter
// 签名
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetAlispSign() string {
    return r.alispSign
}
// Token Setter
// token
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetToken() string {
    return r.token
}
// UserType Setter
// 注册用户类型
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetUserType(userType int64) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

// UserType Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetUserType() int64 {
    return r.userType
}
// AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetAlispTime() string {
    return r.alispTime
}
// Secret Setter
// 一键登录参数
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetSecret(secret string) error {
    r.secret = secret
    r.Set("secret", secret)
    return nil
}

// Secret Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetSecret() string {
    return r.secret
}
// ExtInfo Setter
// json字符串，传入扩展字段
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetExtInfo(extInfo string) error {
    r.extInfo = extInfo
    r.Set("ext_info", extInfo)
    return nil
}

// ExtInfo Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetExtInfo() string {
    return r.extInfo
}
// MtopAppkey Setter
// 选填，调用百川登录接口的appkey，百川登录时，需要传此字段
func (r *AlibabaAlisportsPassportAccountTokenvalidateRequest) SetMtopAppkey(mtopAppkey string) error {
    r.mtopAppkey = mtopAppkey
    r.Set("mtop_appkey", mtopAppkey)
    return nil
}

// MtopAppkey Getter
func (r AlibabaAlisportsPassportAccountTokenvalidateRequest) GetMtopAppkey() string {
    return r.mtopAppkey
}
