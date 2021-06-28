package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
翱象ERP核销回调 APIResponse
alibaba.tcls.aelophy.bill.verificate.callback

翱象ERP核销回调
*/
type AlibabaTclsAelophyBillVerificateCallbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaTclsAelophyBillVerificateCallbackResponse `json:"alibaba_tcls_aelophy_bill_verificate_callback_response,omitempty"` 
    AlibabaTclsAelophyBillVerificateCallbackResponse
}

/* model for simplify = false
type AlibabaTclsAelophyBillVerificateCallbackResponse struct {

    // 处理结果
    
    ApiResult  *struct {
        AlibabaTclsAelophyBillVerificateCallbackApiResult  *AlibabaTclsAelophyBillVerificateCallbackApiResult `json:"alibaba_tcls_aelophy_bill_verificate_callback_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaTclsAelophyBillVerificateCallbackResponse struct {

    // 处理结果
    ApiResult   *AlibabaTclsAelophyBillVerificateCallbackApiResult `json:"api_result,omitempty"`

}
