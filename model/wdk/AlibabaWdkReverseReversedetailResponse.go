package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退款详情 APIResponse
alibaba.wdk.reverse.reversedetail

退款详情
*/
type AlibabaWdkReverseReversedetailAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkReverseReversedetailResponse `json:"alibaba_wdk_reverse_reversedetail_response,omitempty"` 
    AlibabaWdkReverseReversedetailResponse
}

/* model for simplify = false
type AlibabaWdkReverseReversedetailResponse struct {

    // result
    
    Result  *struct {
        ReverseResult  *ReverseResult `json:"reverse_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkReverseReversedetailResponse struct {

    // result
    Result   *ReverseResult `json:"result,omitempty"`

}
