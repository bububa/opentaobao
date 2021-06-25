package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包流量发放-面向支付宝用户 APIResponse
alibaba.aliqin.flow.alipay.publish

用户淘宝流量钱包商家给支付宝用户发放流量
*/
type AlibabaAliqinFlowAlipayPublishAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowAlipayPublishResponse `json:"alibaba_aliqin_flow_alipay_publish_response,omitempty"`
}

type AlibabaAliqinFlowAlipayPublishResponse struct {

    // error
    Error   bool `json:"error,omitempty"`

    // value
    Value   string `json:"value,omitempty"`

    // errorCode
    AlicomErrorCode   string `json:"alicom_error_code,omitempty"`

    // errorMsg
    AlicomErrorMsg   string `json:"alicom_error_msg,omitempty"`

}
