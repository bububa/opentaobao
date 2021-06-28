package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退款详情 APIResponse
alibaba.wdk.reverse.reversedetail

退款详情
*/
type AlibabaWdkReverseReversedetailAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_reverse_reversedetail_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ReverseResult `json:"result,omitempty" xml:"