package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取会员信息 API请求
alibaba.alisports.passport.account.getaccountinfo

获取阿里体育会员信息
*/
type AlibabaAlisportsPassportAccountGetaccountinfoRequest struct {
    model.Params
    // 是否获取详情0否1是 默认0
    _needDetail   int64
    // 当前时间戳
    _alispTime   string
    // 业务appkey
    _alispAppKey   string
    // 业务加密参数
    _alispSign   string
    // 查询类型：1.用户的阿里体育id, 4.用户通过登录生成的sso_token
    _type   int64
    // 要查询的值
    _value   string
    // 决定返回值是否包含扩展字段
    _extInfoType   string
}

// 初始化AlibabaAlisportsPassportAccountGetaccountinfoRequest对象
func NewAlibabaAlisportsPassportAccountGetaccountinfoRequest() *AlibabaAlisportsPassportAccountGetaccountinfoRequest{
    return &AlibabaAlisportsPassportAccountGetaccountinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.getaccountinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NeedDetail Setter
// 是否获取详情0否1是 默认0
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetNeedDetail(_needDetail int64) error {
    r._needDetail = _needDetail
    r.Set("need_detail", _needDetail)
    return nil
}

// NeedDetail Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetNeedDetail() int64 {
    return r._needDetail
}
// AlispTime Setter
// 当前时间戳
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispAppKey Setter
// 业务appkey
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispSign Setter
// 业务加密参数
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetAlispSign() string {
    return r._alispSign
}
// Type Setter
// 查询类型：1.用户的阿里体育id, 4.用户通过登录生成的sso_token
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetType() int64 {
    return r._type
}
// Value Setter
// 要查询的值
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetValue(_value string) error {
    r._value = _value
    r.Set("value", _value)
    return nil
}

// Value Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetValue() string {
    return r._value
}
// ExtInfoType Setter
// 决定返回值是否包含扩展字段
func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetExtInfoType(_extInfoType string) error {
    r._extInfoType = _extInfoType
    r.Set("ext_info_type", _extInfoType)
    return nil
}

// ExtInfoType Getter
func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetExtInfoType() string {
    return r._extInfoType
}