package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
校验支付密码 APIResponse
alibaba.alsc.crm.customer.checkppw

校验支付密码
*/
type AlibabaAlscCrmCustomerCheckppwAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmCustomerCheckppwResponse `json:"alibaba_alsc_crm_customer_checkppw_response,omitempty"`
}

type AlibabaAlscCrmCustomerCheckppwResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
