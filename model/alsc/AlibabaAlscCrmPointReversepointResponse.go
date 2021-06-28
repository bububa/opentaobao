package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
积分消费回退 APIResponse
alibaba.alsc.crm.point.reversepoint

积分消费回退
*/
type AlibabaAlscCrmPointReversepointAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmPointReversepointResponse
}

type AlibabaAlscCrmPointReversepointResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_point_reversepoint_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
