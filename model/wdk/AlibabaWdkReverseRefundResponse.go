package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款打款 APIResponse
alibaba.wdk.reverse.refund

五道口退款打款开放能力接口
*/
type AlibabaWdkReverseRefundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_reverse_refund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaWdkReverseRefundResult `json:"result,omitempty" xml:"