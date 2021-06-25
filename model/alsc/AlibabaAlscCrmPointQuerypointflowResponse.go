package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询积分流水 APIResponse
alibaba.alsc.crm.point.querypointflow

分页查询积分流水
*/
type AlibabaAlscCrmPointQuerypointflowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmPointQuerypointflowResponse `json:"alibaba_alsc_crm_point_querypointflow_response,omitempty"`
}

type AlibabaAlscCrmPointQuerypointflowResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
