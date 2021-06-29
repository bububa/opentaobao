package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
账户充值成功通知 APIRequest
alibaba.fundplatform.account.charge.notify

通知外部业务方充值成功
*/
type AlibabaFundplatformAccountChargeNotifyRequest struct {
    model.Params

    // 入参对象
    request   *AlibabaFundplatformAccountChargeNotifyStruct 

}

func NewAlibabaFundplatformAccountChargeNotifyRequest() *AlibabaFundplatformAccountChargeNotifyRequest{
    return &AlibabaFundplatformAccountChargeNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformAccountChargeNotifyRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.account.charge.notify"
}

func (r AlibabaFundplatformAccountChargeNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformAccountChargeNotifyRequest) SetRequest(request *AlibabaFundplatformAccountChargeNotifyStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r AlibabaFundplatformAccountChargeNotifyRequest) GetRequest() *AlibabaFundplatformAccountChargeNotifyStruct {
    return r.request
}

