package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
回访记录回传API APIResponse
tmall.servicecenter.workcard.call.record

客满回访信息登记
*/
type TmallServicecenterWorkcardCallRecordAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardCallRecordResponse `json:"tmall_servicecenter_workcard_call_record_response,omitempty"`
}

type TmallServicecenterWorkcardCallRecordResponse struct {

    // 返回结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
