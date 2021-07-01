package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelOrderStatusUpdateAPIResponse
订单状态变更 API返回值
alibaba.wdk.channel.order.status.update

订单状态变更 */
type AlibabaWdkChannelOrderStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderStatusUpdateAPIResponseModel
}

// AlibabaWdkChannelOrderStatusUpdateAPIResponseModel is 订单状态变更 成功返回结果
type AlibabaWdkChannelOrderStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderStatusUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
