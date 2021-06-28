package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取顾客优惠券列表 APIResponse
alibaba.alsc.crm.customer.voucher.list

获取顾客优惠券列表
*/
type AlibabaAlscCrmCustomerVoucherListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscCrmCustomerVoucherListResponse `json:"alibaba_alsc_crm_customer_voucher_list_response,omitempty"` 
    AlibabaAlscCrmCustomerVoucherListResponse
}

/* model for simplify = false
type AlibabaAlscCrmCustomerVoucherListResponse struct {

    // 分页返回模型
    
    Result  *struct {
        CommonPageResult  *CommonPageResult `json:"common_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscCrmCustomerVoucherListResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
