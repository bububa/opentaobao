package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
工单完结 APIResponse
tmall.servicecenter.workcard.complete

工单完结
*/
type TmallServicecenterWorkcardCompleteAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardCompleteResponse `json:"tmall_servicecenter_workcard_complete_response,omitempty"`
}

type TmallServicecenterWorkcardCompleteResponse struct {

    // 响应结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
