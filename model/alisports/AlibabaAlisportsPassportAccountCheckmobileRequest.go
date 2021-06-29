package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统--手机号验证接口 API请求
alibaba.alisports.passport.account.checkmobile

验证三方用户的手机号，获取对应的阿里体育会员id
*/
type AlibabaAlisportsPassportAccountCheckmobileRequest struct {
    model.Params
    // 业务appkey
    alispAppKey   string
    // 调用时间戳
    alispTime   string
    // 签名字符串
    alispSign   string
    // 合作方用户ID
    appUid   string
    // 用户呢称
    nick   string
    // 手机号
    mobile   string
    // 性别 0未设置 1男 2女 3保密
    gender   string
    // 生日
    birthday   string
}

// 初始化AlibabaAlisportsPassportAccountCheckmobileRequest对象
func NewAlibabaAlisportsPassportAccountCheckmobileRequest() *AlibabaAlisportsPassportAccountCheckmobileRequest{
    return &AlibabaAlisportsPassportAccountCheckmobileRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.checkmobile"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 业务appkey
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispTime Setter
// 调用时间戳
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetAlispSign() string {
    return r.alispSign
}
// AppUid Setter
// 合作方用户ID
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetAppUid(appUid string) error {
    r.appUid = appUid
    r.Set("app_uid", appUid)
    return nil
}

// AppUid Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetAppUid() string {
    return r.appUid
}
// Nick Setter
// 用户呢称
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetNick() string {
    return r.nick
}
// Mobile Setter
// 手机号
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetMobile() string {
    return r.mobile
}
// Gender Setter
// 性别 0未设置 1男 2女 3保密
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetGender() string {
    return r.gender
}
// Birthday Setter
// 生日
func (r *AlibabaAlisportsPassportAccountCheckmobileRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

// Birthday Getter
func (r AlibabaAlisportsPassportAccountCheckmobileRequest) GetBirthday() string {
    return r.birthday
}
