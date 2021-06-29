package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
移库单获取 APIResponse
alibaba.wdk.ums.shift.get

移库单获取
*/
type AlibabaWdkUmsShiftGetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkUmsShiftGetResponse
}

type AlibabaWdkUmsShiftGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_ums_shift_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *UtmsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
