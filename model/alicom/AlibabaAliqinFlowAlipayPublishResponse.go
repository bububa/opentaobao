package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包流量发放-面向支付宝用户 APIResponse
alibaba.aliqin.flow.alipay.publish

用户淘宝流量钱包商家给支付宝用户发放流量
*/
type AlibabaAliqinFlowAlipayPublishAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_alipay_publish_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // error
    
    Error   bool `json:"error,omitempty" xml:"