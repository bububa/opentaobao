package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建顾客 APIResponse
alibaba.alsc.crm.customer.create

开放本地生活创建顾客功能
*/
type AlibabaAlscCrmCustomerCreateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCustomerCreateResponse
}

type AlibabaAlscCrmCustomerCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_customer_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
