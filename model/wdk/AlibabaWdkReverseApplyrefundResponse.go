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
    AlibabaWdkReverseApplyrefundResponse
}

type AlibabaWdkReverseApplyrefundResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_reverse_applyrefund_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回result
    
    Result   *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
