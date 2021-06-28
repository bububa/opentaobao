package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈鉴定结果 APIResponse
tmall.servicecenter.workcard.evaluate

服务商反馈鉴定结果
*/
type TmallServicecenterWorkcardEvaluateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardEvaluateResponse
}

type TmallServicecenterWorkcardEvaluateResponse struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_evaluate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回值Result
    
    Result   *TmallServicecenterWorkcardEvaluateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
