package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
质量反馈（入库辅助）-ERP下发单 APIResponse
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单
*/
type AlibabaWdkUmsFeedbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkUmsFeedbackResponse `json:"alibaba_wdk_ums_feedback_response,omitempty"` 
    AlibabaWdkUmsFeedbackResponse
}

/* model for simplify = false
type AlibabaWdkUmsFeedbackResponse struct {

    // result
    
    Result  *struct {
        UtmsResult  *UtmsResult `json:"utms_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkUmsFeedbackResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
