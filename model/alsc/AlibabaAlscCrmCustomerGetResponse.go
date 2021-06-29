package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询顾客详情 APIResponse
alibaba.alsc.crm.customer.get

查询顾客详情
*/
type AlibabaAlscCrmCustomerGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCustomerGetResponse
}

type AlibabaAlscCrmCustomerGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_customer_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
