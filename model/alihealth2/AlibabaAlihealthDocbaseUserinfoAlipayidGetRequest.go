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
    _alihealthUserId   string
    // 渠道alipay taobao uc gaode
    _appChannel   string
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
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAlihealthUserId(_alihealthUserId string) error {
    r._alihealthUserId = _alihealthUserId
    r.Set("alihealth_user_id", _alihealthUserId)
    return nil
}

// AlihealthUserId Getter
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetAlihealthUserId() string {
    return r._alihealthUserId
}
// AppChannel Setter
// 渠道alipay taobao uc gaode
func (r *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAppChannel(_appChannel string) error {
    r._appChannel = _appChannel
    r.Set("app_channel", _appChannel)
    return nil
}

// AppChannel Getter
func (r AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) GetAppChannel() string {
    return r._appChannel
}
