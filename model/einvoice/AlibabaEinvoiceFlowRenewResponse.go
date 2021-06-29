package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
工单(入驻、加盘、续约)续约 APIResponse
alibaba.einvoice.flow.renew

工单(含入驻、加盘、续约工单)续约能力开放
*/
type AlibabaEinvoiceFlowRenewAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceFlowRenewResponse
}

type AlibabaEinvoiceFlowRenewResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_flow_renew_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
