package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
共享库存逆向订单理赔单回传 APIResponse
alibaba.wdkorder.sharestock.insurance.refundcallback

共享库存逆向订单理赔单回传
*/
type AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkorderSharestockInsuranceRefundcallbackResponse `json:"alibaba_wdkorder_sharestock_insurance_refundcallback_response,omitempty"` 
    AlibabaWdkorderSharestockInsuranceRefundcallbackResponse
}

/* model for simplify = false
type AlibabaWdkorderSharestockInsuranceRefundcallbackResponse struct {

    // 返回结果
    
    Result  *struct {
        MaochaoOrderInsuranceRefundCallbackResult  *MaochaoOrderInsuranceRefundCallbackResult `json:"maochao_order_insurance_refund_callback_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkorderSharestockInsuranceRefundcallbackResponse struct {

    // 返回结果
    Result   *MaochaoOrderInsuranceRefundCallbackResult `json:"result,omitempty"`

}
