package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改支付密码 APIResponse
alibaba.alsc.crm.customer.updateppw

修改支付密码
*/
type AlibabaAlscCrmCustomerUpdateppwAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCustomerUpdateppwResponse
}

type AlibabaAlscCrmCustomerUpdateppwResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_customer_updateppw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
