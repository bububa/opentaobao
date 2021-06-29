package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取健康id APIRequest
alibaba.alihealth.uic.userinfo.healthid.get

根据支付宝用户ID获取用户健康ID
*/
type AlibabaAlihealthUicUserinfoHealthidGetRequest struct {
    model.Params

    // 支付宝用户ID
    alipayUserId   string 

}

func NewAlibabaAlihealthUicUserinfoHealthidGetRequest() *AlibabaAlihealthUicUserinfoHealthidGetRequest{
    return &AlibabaAlihealthUicUserinfoHealthidGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthUicUserinfoHealthidGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.uic.userinfo.healthid.get"
}

func (r AlibabaAlihealthUicUserinfoHealthidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthUicUserinfoHealthidGetRequest) SetAlipayUserId(alipayUserId string) error {
    r.alipayUserId = alipayUserId
    r.Set("alipay_user_id", alipayUserId)
    return nil
}

func (r AlibabaAlihealthUicUserinfoHealthidGetRequest) GetAlipayUserId() string {
    return r.alipayUserId
}

