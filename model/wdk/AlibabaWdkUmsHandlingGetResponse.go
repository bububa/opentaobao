package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
加工单-回流单（新接口） APIResponse
alibaba.wdk.ums.handling.get

加工单-回流单（新接口）
*/
type AlibabaWdkUmsHandlingGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_ums_handling_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"