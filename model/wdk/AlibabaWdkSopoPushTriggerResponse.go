package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
猫超共享库存寄售sopo推送触发 APIResponse
alibaba.wdk.sopo.push.trigger

猫超共享库存寄售sopo触发推送给商家
*/
type AlibabaWdkSopoPushTriggerAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSopoPushTriggerResponse
}

type AlibabaWdkSopoPushTriggerResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sopo_push_trigger_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 根据站点名称查询产品
    
    Result   *AlibabaWdkSopoPushTriggerApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
