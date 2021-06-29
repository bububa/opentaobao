package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育三方ID绑定接口 API请求
alibaba.alisports.passport.account.bindthirdid

阿里体育三方ID绑定接口
*/
type AlibabaAlisportsPassportAccountBindthirdidRequest struct {
    model.Params
    // 业务方appkey
    _alispAppKey   string
    // 时间戳精确到秒
    _alispTime   string
    // 接口签名
    _alispSign   string
    // 阿里体育用户ID
    _aliuid   string
    // 三方ID
    _appUid   string
    // 手机号
    _mobile   string
}

// 初始化AlibabaAlisportsPassportAccountBindthirdidRequest对象
func NewAlibabaAlisportsPassportAccountBindthirdidRequest() *AlibabaAlisportsPassportAccountBindthirdidRequest{
    return &AlibabaAlisportsPassportAccountBindthirdidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.bindthirdid"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 业务方appkey
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAlispSign() string {
    return r._alispSign
}
// Aliuid Setter
// 阿里体育用户ID
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAliuid() string {
    return r._aliuid
}
// AppUid Setter
// 三方ID
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetAppUid(_appUid string) error {
    r._appUid = _appUid
    r.Set("app_uid", _appUid)
    return nil
}

// AppUid Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetAppUid() string {
    return r._appUid
}
// Mobile Setter
// 手机号
func (r *AlibabaAlisportsPassportAccountBindthirdidRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaAlisportsPassportAccountBindthirdidRequest) GetMobile() string {
    return r._mobile
}
