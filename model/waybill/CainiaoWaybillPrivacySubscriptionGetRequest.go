package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订购查询 API请求
cainiao.waybill.privacy.subscription.get

ISV查询商家是否订购隐私面单
*/
type CainiaoWaybillPrivacySubscriptionGetRequest struct {
    model.Params
}

// 初始化CainiaoWaybillPrivacySubscriptionGetRequest对象
func NewCainiaoWaybillPrivacySubscriptionGetRequest() *CainiaoWaybillPrivacySubscriptionGetRequest{
    return &CainiaoWaybillPrivacySubscriptionGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillPrivacySubscriptionGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.privacy.subscription.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillPrivacySubscriptionGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
