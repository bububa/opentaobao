package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务平台工单取消接口 APIResponse
alibaba.servicecenter.workcard.cancel

取消服务工单
*/
type AlibabaServicecenterWorkcardCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_servicecenter_workcard_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回参数
    
    Result   *FulfilplatformResult `json:"result,omitempty" xml:"