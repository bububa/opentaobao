package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelorderrefundconfirmAPIResponse 退款确认 API返回值
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
type AlibabawdkchannelorderrefundconfirmAPIResponse struct {
	model.CommonResponse
	AlibabawdkchannelorderrefundconfirmAPIResponseModel
}

// AlibabawdkchannelorderrefundconfirmAPIResponseModel is 退款确认 成功返回结果
type AlibabawdkchannelorderrefundconfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_refund_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabawdkchannelorderrefundconfirmApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
