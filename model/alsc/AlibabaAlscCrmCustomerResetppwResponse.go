package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
重置支付密码 APIResponse
alibaba.alsc.crm.customer.resetppw

重置支付密码
*/
type AlibabaAlscCrmCustomerResetppwAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmCustomerResetppwResponse `json:"alibaba_alsc_crm_customer_resetppw_response,omitempty"`
}

type AlibabaAlscCrmCustomerResetppwResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
