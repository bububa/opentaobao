package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口外部商户老pos机产生的退款单同步进盒马 APIResponse
alibaba.wdk.oldpos.refund.create

淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
*/
type AlibabaWdkOldposRefundCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOldposRefundCreateResponse `json:"alibaba_wdk_oldpos_refund_create_response,omitempty"` 
    AlibabaWdkOldposRefundCreateResponse
}

/* model for simplify = false
type AlibabaWdkOldposRefundCreateResponse struct {

    // result
    
    Result  *struct {
        PosRefundCreateResult  *PosRefundCreateResult `json:"pos_refund_create_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOldposRefundCreateResponse struct {

    // result
    Result   *PosRefundCreateResult `json:"result,omitempty"`

}
