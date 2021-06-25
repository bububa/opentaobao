package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询顾客详情 APIResponse
alibaba.alsc.crm.customer.get

查询顾客详情
*/
type AlibabaAlscCrmCustomerGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmCustomerGetResponse `json:"alibaba_alsc_crm_customer_get_response,omitempty"`
}

type AlibabaAlscCrmCustomerGetResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
