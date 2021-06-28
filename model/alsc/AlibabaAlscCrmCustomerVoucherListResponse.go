package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取顾客优惠券列表 APIResponse
alibaba.alsc.crm.customer.voucher.list

获取顾客优惠券列表
*/
type AlibabaAlscCrmCustomerVoucherListAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCustomerVoucherListResponse
}

type AlibabaAlscCrmCustomerVoucherListResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_customer_voucher_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
