package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractactivityunregisterAPIResponse ISV互动应用活动关闭服务 API返回值
// alibaba.interact.activity.unregister
//
// 卖家在ISV互动应用中设置的活动主动关闭的服务
type AlibabainteractactivityunregisterAPIResponse struct {
	model.CommonResponse
	AlibabainteractactivityunregisterAPIResponseModel
}

// AlibabainteractactivityunregisterAPIResponseModel is ISV互动应用活动关闭服务 成功返回结果
type AlibabainteractactivityunregisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_activity_unregister_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关闭活动结果
	UnregisterResult *AllsparkResult `json:"unregister_result,omitempty" xml:"unregister_result,omitempty"`
}
