package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
加工单-回流单（新接口） APIResponse
alibaba.wdk.ums.handling.get

加工单-回流单（新接口）
*/
type AlibabaWdkUmsHandlingGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkUmsHandlingGetResponse `json:"alibaba_wdk_ums_handling_get_response,omitempty"`
}

type AlibabaWdkUmsHandlingGetResponse struct {

    // result
    Result   *UtmsResult `json:"result,omitempty"`

}
