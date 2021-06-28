package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建顾客 APIResponse
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能
*/
type AlibabaAlscCrmCustomerCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCustomerCreateResponse `json:"alibaba_alsc_crm_customer_create_response,omitempty"` 
    AlibabaAlscCrmCustomerCreateResponse
}

/* model for simplify = false
type AlibabaAlscCrmCustomerCreateResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCustomerCreateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
