package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈鉴定结果 APIResponse
tmall.servicecenter.workcard.evaluate

服务商反馈鉴定结果
*/
type TmallServicecenterWorkcardEvaluateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardEvaluateResponse `json:"tmall_servicecenter_workcard_evaluate_response,omitempty"`
}

type TmallServicecenterWorkcardEvaluateResponse struct {

    // 返回值Result
    Result   *TmallServicecenterWorkcardEvaluateResult `json:"result,omitempty"`

}
