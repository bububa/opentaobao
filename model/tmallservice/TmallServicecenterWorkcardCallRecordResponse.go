package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回访记录回传API APIResponse
tmall.servicecenter.workcard.call.record

客满回访信息登记
*/
type TmallServicecenterWorkcardCallRecordAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_servicecenter_workcard_call_record_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"