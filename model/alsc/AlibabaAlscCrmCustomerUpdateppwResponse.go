package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改支付密码 APIResponse
alibaba.alsc.crm.customer.updateppw

修改支付密码
*/
type AlibabaAlscCrmCustomerUpdateppwAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmCustomerUpdateppwResponse `json:"alibaba_alsc_crm_customer_updateppw_response,omitempty"`
}

type AlibabaAlscCrmCustomerUpdateppwResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
