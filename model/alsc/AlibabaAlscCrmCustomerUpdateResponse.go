package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新顾客信息 APIResponse
alibaba.alsc.crm.customer.update

更新顾客信息
*/
type AlibabaAlscCrmCustomerUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCustomerUpdateResponse
}

type AlibabaAlscCrmCustomerUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_customer_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
