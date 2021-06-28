package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
移库单获取 APIResponse
alibaba.wdk.ums.shift.get

移库单获取
*/
type AlibabaWdkUmsShiftGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkUmsShiftGetResponse `json:"alibaba_wdk_ums_shift_get_response,omitempty"` 
    AlibabaWdkUmsShiftGetResponse
}

/* model for simplify = false
type AlibabaWdkUmsShiftGetResponse struct {

    // result
    
    Result  *struct {
        UtmsResult  *UtmsResult `json:"utms_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkUmsShiftGetResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
