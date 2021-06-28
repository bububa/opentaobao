package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
保存和更新顾客 APIResponse
alibaba.alsc.crm.open.customer.save

用来保存顾客，如果已经存在的话，则更新顾客
*/
type AlibabaAlscCrmOpenCustomerSaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmOpenCustomerSaveResponse `json:"alibaba_alsc_crm_open_customer_save_response,omitempty"` 
    AlibabaAlscCrmOpenCustomerSaveResponse
}

/* model for simplify = false
type AlibabaAlscCrmOpenCustomerSaveResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmOpenCustomerSaveResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
