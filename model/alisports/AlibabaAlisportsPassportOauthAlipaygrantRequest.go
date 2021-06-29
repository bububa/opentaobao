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
    _alispAppKey   string
    // 请求接口的时间戳
    _alispTime   string
    // 签名字符串
    _alispSign   string
    // 调用支付宝登录sdk返回的code
    _authCode   string
    // 固定为rich_sports
    _partnerMode   string
    // 支付宝的appid
    _appid   string
    // 合作方的userid，即用户唯一的id标识
    _appUid   string
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
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispTime Setter
// 请求接口的时间戳
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAlispSign() string {
    return r._alispSign
}
// AuthCode Setter
// 调用支付宝登录sdk返回的code
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAuthCode(_authCode string) error {
    r._authCode = _authCode
    r.Set("auth_code", _authCode)
    return nil
}

// AuthCode Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAuthCode() string {
    return r._authCode
}
// PartnerMode Setter
// 固定为rich_sports
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetPartnerMode(_partnerMode string) error {
    r._partnerMode = _partnerMode
    r.Set("partner_mode", _partnerMode)
    return nil
}

// PartnerMode Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetPartnerMode() string {
    return r._partnerMode
}
// Appid Setter
// 支付宝的appid
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAppid(_appid string) error {
    r._appid = _appid
    r.Set("appid", _appid)
    return nil
}

// Appid Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAppid() string {
    return r._appid
}
// AppUid Setter
// 合作方的userid，即用户唯一的id标识
func (r *AlibabaAlisportsPassportOauthAlipaygrantRequest) SetAppUid(_appUid string) error {
    r._appUid = _appUid
    r.Set("app_uid", _appUid)
    return nil
}

// AppUid Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantRequest) GetAppUid() string {
    return r._appUid
}
