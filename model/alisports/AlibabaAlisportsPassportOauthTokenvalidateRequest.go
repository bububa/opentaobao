package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统--获取登录信息接口 API请求
alibaba.alisports.passport.oauth.tokenvalidate

阿里体育会员系统--获取登录信息接口
*/
type AlibabaAlisportsPassportOauthTokenvalidateRequest struct {
    model.Params
    // 登录时返回给前端的token
    _token   string
    // 时间戳
    _alispTime   string
    // 应用的appkey
    _alispAppKey   string
    // 参数加密之后的串
    _alispSign   string
}

// 初始化AlibabaAlisportsPassportOauthTokenvalidateRequest对象
func NewAlibabaAlisportsPassportOauthTokenvalidateRequest() *AlibabaAlisportsPassportOauthTokenvalidateRequest{
    return &AlibabaAlisportsPassportOauthTokenvalidateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.oauth.tokenvalidate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 登录时返回给前端的token
func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetToken() string {
    return r._token
}
// AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispAppKey Setter
// 应用的appkey
func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispSign Setter
// 参数加密之后的串
func (r *AlibabaAlisportsPassportOauthTokenvalidateRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportOauthTokenvalidateRequest) GetAlispSign() string {
    return r._alispSign
}
