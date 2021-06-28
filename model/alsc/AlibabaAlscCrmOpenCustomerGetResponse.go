package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询会员资产 APIResponse
alibaba.alsc.crm.open.customer.get

查询会员身份，资产等
*/
type AlibabaAlscCrmOpenCustomerGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmOpenCustomerGetResponse `json:"alibaba_alsc_crm_open_customer_get_response,omitempty"` 
    AlibabaAlscCrmOpenCustomerGetResponse
}

/* model for simplify = false
type AlibabaAlscCrmOpenCustomerGetResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmOpenCustomerGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
