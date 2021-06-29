package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取会员信息 APIRequest
alibaba.alisports.passport.account.getaccountinfo

获取阿里体育会员信息
*/
type AlibabaAlisportsPassportAccountGetaccountinfoRequest struct {
    model.Params

    // 是否获取详情0否1是 默认0
    needDetail   int64 

    // 当前时间戳
    alispTime   string 

    // 业务appkey
    alispAppKey   string 

    // 业务加密参数
    alispSign   string 

    // 查询类型：1.用户的阿里体育id, 4.用户通过登录生成的sso_token
    type   int64 

    // 要查询的值
    value   string 

    // 决定返回值是否包含扩展字段
    extInfoType   string 

}

func NewAlibabaAlisportsPassportAccountGetaccountinfoRequest() *AlibabaAlisportsPassportAccountGetaccountinfoRequest{
    return &AlibabaAlisportsPassportAccountGetaccountinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.getaccountinfo"
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetNeedDetail(needDetail int64) error {
    r.needDetail = needDetail
    r.Set("need_detail", needDetail)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetNeedDetail() int64 {
    return r.needDetail
}

func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetValue(value string) error {
    r.value = value
    r.Set("value", value)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetValue() string {
    return r.value
}

func (r *AlibabaAlisportsPassportAccountGetaccountinfoRequest) SetExtInfoType(extInfoType string) error {
    r.extInfoType = extInfoType
    r.Set("ext_info_type", extInfoType)
    return nil
}

func (r AlibabaAlisportsPassportAccountGetaccountinfoRequest) GetExtInfoType() string {
    return r.extInfoType
}

