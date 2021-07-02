package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSopoPushTriggerAPIResponse 猫超共享库存寄售sopo推送触发 API返回值
// alibaba.wdk.sopo.push.trigger
//
// 猫超共享库存寄售sopo触发推送给商家
type AlibabaWdkSopoPushTriggerAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSopoPushTriggerAPIResponseModel
}

// AlibabaWdkSopoPushTriggerAPIResponseModel is 猫超共享库存寄售sopo推送触发 成功返回结果
type AlibabaWdkSopoPushTriggerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sopo_push_trigger_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *AlibabaWdkSopoPushTriggerApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
