package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据健康ID获取支付宝ID APIRequest
alibaba.alihealth.docbase.userinfo.alipayid.get

根据健康ID获取支付宝ID
*/
type AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest struct {
    model.Params

    // 阿里健康ID
    alihealthUserId   string 

    // 渠道alipay taobao uc gaode
    appChannel   string 

}

func NewAlibabaAlihealthDocbaseUserinfoAlipayidGetRequest() *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest{
    return &AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.docbase.userinfo.alipayid.get"
}

func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAlihealthUserId(alihealthUserId string) error {
    r.alihealthUserId = alihealthUserId
    r.Set("alihealth_user_id", alihealthUserId)
    return nil
}

func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetAlihealthUserId() string {
    return r.alihealthUserId
}

func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAppChannel(appChannel string) error {
    r.appChannel = appChannel
    r.Set("app_channel", appChannel)
    return nil
}

func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetAppChannel() string {
    return r.appChannel
}

