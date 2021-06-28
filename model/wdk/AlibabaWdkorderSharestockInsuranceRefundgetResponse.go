package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
共享库存投保业务售后逆向订单数据获取 APIResponse
alibaba.wdkorder.sharestock.insurance.refundget

共享库存投保业务售后逆向订单数据获取
*/
type AlibabaWdkorderSharestockInsuranceRefundgetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkorderSharestockInsuranceRefundgetResponse `json:"alibaba_wdkorder_sharestock_insurance_refundget_response,omitempty"` 
    AlibabaWdkorderSharestockInsuranceRefundgetResponse
}

/* model for simplify = false
type AlibabaWdkorderSharestockInsuranceRefundgetResponse struct {

    // 返回结果
    
    Result  *struct {
        MaochaoOrderInsuranceRefundQueryResult  *MaochaoOrderInsuranceRefundQueryResult `json:"maochao_order_insurance_refund_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkorderSharestockInsuranceRefundgetResponse struct {

    // 返回结果
    Result   *MaochaoOrderInsuranceRefundQueryResult `json:"result,omitempty"`

}
