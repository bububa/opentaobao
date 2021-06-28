package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存逆向订单理赔单回传 APIResponse
alibaba.wdkorder.sharestock.insurance.refundcallback

共享库存逆向订单理赔单回传
*/
type AlibabaWdkorderSharestockInsuranceRefundcallbackAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdkorder_sharestock_insurance_refundcallback_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *MaochaoOrderInsuranceRefundCallbackResult `json:"result,omitempty" xml:"