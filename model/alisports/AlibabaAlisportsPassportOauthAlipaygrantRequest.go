package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统-支付宝授权接口 API请求
alibaba.alisports.passport.oauth.alipaygrant

开放给乐心运动使用的支付宝授权接口
*/
type AlibabaAlisportsPassportOauthAlipaygrantRequest struct {
    model.Params
    // 阿里体育分配的用户appkey
    alispAppKey   string
    // 请求接口的时间戳
    alispTime   string
    // 签名字符串
    alispSign   string
    // 调用支付宝登录sdk返回的code
    authCode   string
    // 固定为rich_sports
    partnerMode   string
    // 支付宝的appid
    appid   string
    // 合作方的userid，即用户唯一的id标识
    appUid   string
}

// 初始化AlibabaAlisportsPassportOauthAlipaygrantRequest对象
func NewAlibabaAlisportsPassportOauthAlipaygrantRequest() *AlibabaAlisportsPassportOauthAlipaygrantRequest{
    return &AlibabaAlisportsPassportOauthAlipaygrantRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.oauth.alipaygrant"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 阿里体育分配的用户appkey
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispTime Setter
// 请求接口的时间戳
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAlispSign() string {
    return r.alispSign
}
// AuthCode Setter
// 调用支付宝登录sdk返回的code
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAuthCode(authCode string) error {
    r.authCode = authCode
    r.Set("auth_code", authCode)
    return nil
}

// AuthCode Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAuthCode() string {
    return r.authCode
}
// PartnerMode Setter
// 固定为rich_sports
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetPartnerMode(partnerMode string) error {
    r.partnerMode = partnerMode
    r.Set("partner_mode", partnerMode)
    return nil
}

// PartnerMode Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetPartnerMode() string {
    return r.partnerMode
}
// Appid Setter
// 支付宝的appid
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAppid(appid string) error {
    r.appid = appid
    r.Set("appid", appid)
    return nil
}

// Appid Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAppid() string {
    return r.appid
}
// AppUid Setter
// 合作方的userid，即用户唯一的id标识
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAppUid(appUid string) error {
    r.appUid = appUid
    r.Set("app_uid", appUid)
    return nil
}

// AppUid Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAppUid() string {
    return r.appUid
}
