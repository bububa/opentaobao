package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈鉴定结果 API返回值 
tmall.servicecenter.workcard.evaluate

服务商反馈鉴定结果
*/
type TmallServicecenterWorkcardEvaluateAPIResponse struct {
    model.CommonResponse
    TmallServicecenterWorkcardEvaluateAPIResponseModel
}

// 服务商反馈鉴定结果 成功返回结果
type TmallServicecenterWorkcardEvaluateAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_servicecenter_workcard_evaluate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回值Result
    Result   *TmallServicecenterWorkcardEvaluateResult `json:"result,omitempty" xml:"result,omitempty"`
}
