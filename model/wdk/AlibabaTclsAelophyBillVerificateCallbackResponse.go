package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象ERP核销回调 APIResponse
alibaba.tcls.aelophy.bill.verificate.callback

翱象ERP核销回调
*/
type AlibabaTclsAelophyBillVerificateCallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_tcls_aelophy_bill_verificate_callback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 处理结果
    
    ApiResult   *AlibabaTclsAelophyBillVerificateCallbackApiResult `json:"api_result,omitempty" xml:"