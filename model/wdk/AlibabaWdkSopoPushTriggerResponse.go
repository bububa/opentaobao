package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
猫超共享库存寄售sopo推送触发 APIResponse
alibaba.wdk.sopo.push.trigger

猫超共享库存寄售sopo触发推送给商家
*/
type AlibabaWdkSopoPushTriggerAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSopoPushTriggerResponse `json:"alibaba_wdk_sopo_push_trigger_response,omitempty"`
}

type AlibabaWdkSopoPushTriggerResponse struct {

    // 根据站点名称查询产品
    Result   *AlibabaWdkSopoPushTriggerApiResult `json:"result,omitempty"`

}
