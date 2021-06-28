package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
质量反馈（入库辅助）-ERP下发单 APIResponse
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单
*/
type AlibabaWdkUmsFeedbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_feedback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"