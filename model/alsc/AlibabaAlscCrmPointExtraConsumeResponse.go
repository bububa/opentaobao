package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
积分补扣 APIResponse
alibaba.alsc.crm.point.extra.consume

积分补扣
*/
type AlibabaAlscCrmPointExtraConsumeAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmPointExtraConsumeResponse
}

type AlibabaAlscCrmPointExtraConsumeResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_point_extra_consume_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
