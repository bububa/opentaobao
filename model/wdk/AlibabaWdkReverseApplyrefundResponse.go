package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向申请接口 APIResponse
alibaba.wdk.reverse.applyrefund

逆向渲染
*/
type AlibabaWdkReverseApplyrefundAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_reverse_applyrefund_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回result
    
    Result   *ReverseResult `json:"result,omitempty" xml:"