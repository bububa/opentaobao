package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员资产 APIResponse
alibaba.alsc.crm.open.customer.get

查询会员身份，资产等
*/
type AlibabaAlscCrmOpenCustomerGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmOpenCustomerGetResponse
}

type AlibabaAlscCrmOpenCustomerGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_open_customer_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
