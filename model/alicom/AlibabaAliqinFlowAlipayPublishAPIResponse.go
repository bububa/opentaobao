package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowAlipayPublishAPIResponse 流量钱包流量发放-面向支付宝用户 API返回值
// alibaba.aliqin.flow.alipay.publish
//
// 用户淘宝流量钱包商家给支付宝用户发放流量
type AlibabaAliqinFlowAlipayPublishAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFlowAlipayPublishAPIResponseModel
}

// AlibabaAliqinFlowAlipayPublishAPIResponseModel is 流量钱包流量发放-面向支付宝用户 成功返回结果
type AlibabaAliqinFlowAlipayPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_alipay_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// errorCode
	AlicomErrorCode string `json:"alicom_error_code,omitempty" xml:"alicom_error_code,omitempty"`
	// errorMsg
	AlicomErrorMsg string `json:"alicom_error_msg,omitempty" xml:"alicom_error_msg,omitempty"`
}
