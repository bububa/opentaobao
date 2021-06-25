package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
逆向提交 APIResponse
alibaba.wdk.reverse.creatrefund

逆向申请提交
*/
type AlibabaWdkReverseCreatrefundAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkReverseCreatrefundResponse `json:"alibaba_wdk_reverse_creatrefund_response,omitempty"`
}

type AlibabaWdkReverseCreatrefundResponse struct {

    // ReverseResult
    Result   *ReverseResult `json:"result,omitempty"`

}
