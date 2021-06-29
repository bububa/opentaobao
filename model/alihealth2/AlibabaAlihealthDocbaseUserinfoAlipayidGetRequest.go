package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据健康ID获取支付宝ID API请求
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

// 初始化AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest对象
func NewAlibabaAlihealthDocbaseUserinfoAlipayidGetRequest() *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest{
    return &AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.docbase.userinfo.alipayid.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlihealthUserId Setter
// 阿里健康ID
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAlihealthUserId(alihealthUserId string) error {
    r.alihealthUserId = alihealthUserId
    r.Set("alihealth_user_id", alihealthUserId)
    return nil
}

// AlihealthUserId Getter
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetAlihealthUserId() string {
    return r.alihealthUserId
}
// AppChannel Setter
// 渠道alipay taobao uc gaode
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAppChannel(appChannel string) error {
    r.appChannel = appChannel
    r.Set("app_channel", appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetAppChannel() string {
    return r.appChannel
}
