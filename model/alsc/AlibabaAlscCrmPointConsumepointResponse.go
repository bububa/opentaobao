package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
积分抵现 APIResponse
alibaba.alsc.crm.point.consumepoint

积分抵现
*/
type AlibabaAlscCrmPointConsumepointAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmPointConsumepointResponse `json:"alibaba_alsc_crm_point_consumepoint_response,omitempty"`
}

type AlibabaAlscCrmPointConsumepointResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
