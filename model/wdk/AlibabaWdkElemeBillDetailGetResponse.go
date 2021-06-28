package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
饿了么对账单查询，带订单明细 APIResponse
alibaba.wdk.eleme.bill.detail.get

查询饿了么对账单信息，带订单明细
*/
type AlibabaWdkElemeBillDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkElemeBillDetailGetResponse `json:"alibaba_wdk_eleme_bill_detail_get_response,omitempty"` 
    AlibabaWdkElemeBillDetailGetResponse
}

/* model for simplify = false
type AlibabaWdkElemeBillDetailGetResponse struct {

    // 返回结果
    
    Result  *struct {
        AlibabaWdkElemeBillDetailGetApiResult  *AlibabaWdkElemeBillDetailGetApiResult `json:"alibaba_wdk_eleme_bill_detail_get_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkElemeBillDetailGetResponse struct {

    // 返回结果
    Result   *AlibabaWdkElemeBillDetailGetApiResult `json:"result,omitempty"`

}
