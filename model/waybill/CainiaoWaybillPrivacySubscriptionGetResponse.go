package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
隐私面单商家订购查询 APIResponse
cainiao.waybill.privacy.subscription.get

ISV查询商家是否订购隐私面单
*/
type CainiaoWaybillPrivacySubscriptionGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"cainiao_waybill_privacy_subscription_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *CainiaoWaybillPrivacySubscriptionGetResult `json:"result,omitempty" xml:"