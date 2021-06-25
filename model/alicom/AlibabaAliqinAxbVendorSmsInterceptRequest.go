package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
AXB短信托收推送接口 APIRequest
alibaba.aliqin.axb.vendor.sms.intercept

用于给供应商推送需要托收的短信
*/
type AlibabaAliqinAxbVendorSmsInterceptRequest struct {
    model.Params

    // 短信托收结构体
    smsInterceptRequest   *SmsInterceptRequest 

}

func NewAlibabaAliqinAxbVendorSmsInterceptRequest() *AlibabaAliqinAxbVendorSmsInterceptRequest{
    return &AlibabaAliqinAxbVendorSmsInterceptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinAxbVendorSmsInterceptRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.sms.intercept"
}

func (r AlibabaAliqinAxbVendorSmsInterceptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinAxbVendorSmsInterceptRequest) SetSmsInterceptRequest(smsInterceptRequest *SmsInterceptRequest) error {
    r.smsInterceptRequest = smsInterceptRequest
    r.Set("sms_intercept_request", smsInterceptRequest)
    return nil
}

func (r AlibabaAliqinAxbVendorSmsInterceptRequest) GetSmsInterceptRequest() *SmsInterceptRequest {
    return r.smsInterceptRequest
}

