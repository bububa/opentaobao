package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商户取消 APIResponse
alibaba.ele.fengniao.cancel.merchant

商户取消配送
*/
type AlibabaEleFengniaoCancelMerchantAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_ele_fengniao_cancel_merchant_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // message
    
    Message   string `json:"message,omitempty" xml:"