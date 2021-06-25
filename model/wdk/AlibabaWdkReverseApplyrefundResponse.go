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
    Response *AlibabaWdkReverseApplyrefundResponse `json:"alibaba_wdk_reverse_applyrefund_response,omitempty"`
}

type AlibabaWdkReverseApplyrefundResponse struct {

    // 接口返回result
    Result   *ReverseResult `json:"result,omitempty"`

}
