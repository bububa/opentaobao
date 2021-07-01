package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderCreateAPIResponse
创建订单 API返回值
alibaba.wdk.channel.order.create

外部商家创建订单 */
type AlibabaWdkChannelOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderCreateAPIResponseModel
}

// AlibabaWdkChannelOrderCreateAPIResponseModel is 创建订单 成功返回结果
type AlibabaWdkChannelOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
