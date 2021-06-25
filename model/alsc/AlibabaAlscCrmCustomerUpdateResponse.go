package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新顾客信息 APIResponse
alibaba.alsc.crm.customer.update

更新顾客信息
*/
type AlibabaAlscCrmCustomerUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmCustomerUpdateResponse `json:"alibaba_alsc_crm_customer_update_response,omitempty"`
}

type AlibabaAlscCrmCustomerUpdateResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
