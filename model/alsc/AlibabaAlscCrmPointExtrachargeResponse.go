package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
积分补录 APIResponse
alibaba.alsc.crm.point.extracharge

积分补录
*/
type AlibabaAlscCrmPointExtrachargeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_point_extracharge_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"