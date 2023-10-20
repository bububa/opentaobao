package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeactivitysubscriptionbindAPIResponse 销售活动绑定认购商品 API返回值
// alibaba.alihouse.newhome.activity.subscription.bind
//
// 销售活动绑定认购商品
type AlibabaalihousenewhomeactivitysubscriptionbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeactivitysubscriptionbindAPIResponseModel
}

// AlibabaalihousenewhomeactivitysubscriptionbindAPIResponseModel is 销售活动绑定认购商品 成功返回结果
type AlibabaalihousenewhomeactivitysubscriptionbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_activity_subscription_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回报文参数
	Result *AlibabaalihousenewhomeactivitysubscriptionbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
