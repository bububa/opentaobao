package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育会员系统--取消三方关联接口 APIRequest
alibaba.alisports.passport.account.delrelation

阿里体育会员系统--取消三方关联接口
*/
type AlibabaAlisportsPassportAccountDelrelationRequest struct {
    model.Params

    // 业务appkey
    alispAppKey   string 

    // 调用时间戳
    alispTime   string 

    // 签名字符串
    alispSign   string 

    // 合作方用户ID
    appUid   string 

    // 阿里体育会员id
    aliuid   string 

}

func NewAlibabaAlisportsPassportAccountDelrelationRequest() *AlibabaAlisportsPassportAccountDelrelationRequest{
    return &AlibabaAlisportsPassportAccountDelrelationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.account.delrelation"
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportAccountDelrelationRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportAccountDelrelationRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportAccountDelrelationRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsPassportAccountDelrelationRequest) SetAppUid(appUid string) error {
    r.appUid = appUid
    r.Set("app_uid", appUid)
    return nil
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetAppUid() string {
    return r.appUid
}

func (r *AlibabaAlisportsPassportAccountDelrelationRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsPassportAccountDelrelationRequest) GetAliuid() string {
    return r.aliuid
}

