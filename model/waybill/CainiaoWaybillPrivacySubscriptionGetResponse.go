package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订购查询 APIResponse
cainiao.waybill.privacy.subscription.get

ISV查询商家是否订购隐私面单
*/
type CainiaoWaybillPrivacySubscriptionGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillPrivacySubscriptionGetResponse `json:"cainiao_waybill_privacy_subscription_get_response,omitempty"` 
    CainiaoWaybillPrivacySubscriptionGetResponse
}

/* model for simplify = false
type CainiaoWaybillPrivacySubscriptionGetResponse struct {

    // 接口返回model
    
    Result  *struct {
        CainiaoWaybillPrivacySubscriptionGetResult  *CainiaoWaybillPrivacySubscriptionGetResult `json:"cainiao_waybill_privacy_subscription_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoWaybillPrivacySubscriptionGetResponse struct {

    // 接口返回model
    Result   *CainiaoWaybillPrivacySubscriptionGetResult `json:"result,omitempty"`

}
