package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订购查询 APIRequest
cainiao.waybill.privacy.subscription.get

ISV查询商家是否订购隐私面单
*/
type CainiaoWaybillPrivacySubscriptionGetRequest struct {
    model.Params

}

func NewCainiaoWaybillPrivacySubscriptionGetRequest() *CainiaoWaybillPrivacySubscriptionGetRequest{
    return &CainiaoWaybillPrivacySubscriptionGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillPrivacySubscriptionGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.privacy.subscription.get"
}

func (r CainiaoWaybillPrivacySubscriptionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


