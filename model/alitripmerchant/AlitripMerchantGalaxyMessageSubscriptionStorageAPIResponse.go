package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymessagesubscriptionstorageAPIResponse 存储模版ID API返回值
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
type AlitripmerchantgalaxymessagesubscriptionstorageAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxymessagesubscriptionstorageAPIResponseModel
}

// AlitripmerchantgalaxymessagesubscriptionstorageAPIResponseModel is 存储模版ID 成功返回结果
type AlitripmerchantgalaxymessagesubscriptionstorageAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_message_subscription_storage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxymessagesubscriptionstorageResponse `json:"result,omitempty" xml:"result,omitempty"`
}
