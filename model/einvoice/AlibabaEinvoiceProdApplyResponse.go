package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发票申请 APIResponse
alibaba.einvoice.prod.apply

提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
*/
type AlibabaEinvoiceProdApplyAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceProdApplyResponse
}

type AlibabaEinvoiceProdApplyResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_prod_apply_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
