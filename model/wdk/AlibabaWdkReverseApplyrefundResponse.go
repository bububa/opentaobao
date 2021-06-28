package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
逆向申请接口 APIResponse
alibaba.wdk.reverse.applyrefund

逆向渲染
*/
type AlibabaWdkReverseApplyrefundAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkReverseApplyrefundResponse `json:"alibaba_wdk_reverse_applyrefund_response,omitempty"` 
    AlibabaWdkReverseApplyrefundResponse
}

/* model for simplify = false
type AlibabaWdkReverseApplyrefundResponse struct {

    // 接口返回result
    
    Result  *struct {
        ReverseResult  *ReverseResult `json:"reverse_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkReverseApplyrefundResponse struct {

    // 接口返回result
    Result   *ReverseResult `json:"result,omitempty"`

}
